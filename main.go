package main

import (
	"log"
	"todo-app/controller"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.POST("/addTodo", controller.AddTodo)
	e.GET("/getTodos", controller.GetTodos)
	e.DELETE("/deleteTodo/:id", controller.DeleteTodo)
	e.PUT("/updateTodo/:id", controller.UpdateTodo)

	log.Fatal(e.Start(":8080"))

}
