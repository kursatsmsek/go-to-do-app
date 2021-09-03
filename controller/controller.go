package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"todo-app/configuration"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Title     string `json:"title"`
	Completed *bool  `json:"completed"`
}

func getDBUri() string {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return "mongodb://localhost:27017"
	}
	return uri
}

func getDBName() string {
	return "todo-service"
}

func AddTodo(c echo.Context) error {

	config := new(configuration.MongoConfiguration).Init(getDBUri(), getDBName())
	collection := config.Database().Collection("todos")

	todo := Todo{}
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &todo)

	if todo.Completed == nil || todo.Title == "" {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	_, err = collection.InsertOne(context.TODO(), bson.D{
		{Key: "Title", Value: todo.Title},
		{Key: "Completed", Value: todo.Completed},
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success")
}

func GetTodos(c echo.Context) error {
	config := new(configuration.MongoConfiguration).Init(getDBUri(), getDBName())
	collection := config.Database().Collection("todos")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	var todos []bson.M
	if err = cursor.All(context.TODO(), &todos); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {

	config := new(configuration.MongoConfiguration).Init(getDBUri(), getDBName())
	collection := config.Database().Collection("todos")
	id := c.Param("id")

	idPrimitive, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": idPrimitive})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, res)

}

func UpdateTodo(c echo.Context) error {
	config := new(configuration.MongoConfiguration).Init(getDBUri(), getDBName())
	collection := config.Database().Collection("todos")

	todo := Todo{}
	id := c.Param("id")
	idPrimitive, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": idPrimitive}

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	json.Unmarshal(body, &todo)

	if todo.Title == "" || todo.Completed == nil {
		return c.JSON(http.StatusBadRequest, "completed and id cannot be null")
	}

	res, err := collection.UpdateOne(context.TODO(), filter, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "Title", Value: todo.Title},
			{Key: "Completed", Value: todo.Completed},
		}},
	})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return c.JSON(http.StatusNoContent, "no match")
	}

	return c.JSON(http.StatusOK, res)

}
