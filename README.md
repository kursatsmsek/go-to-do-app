# Todo App Backend (Golang)

## Run
```
$ docker-compose up -d
```

## Endpoints

### Add Todo - Post
```
http://localhost:8080/addTodo
```
##### Body

``` json 
{
    "Title":"buy a chocolate",
    "Completed":false
}
```
<hr/>

### Update Todo - Put
```
http://localhost:8080/updateTodo/id
```
##### Body

``` json 
{
    "Title":"buy a chocolate",
    "Completed":false
}
```
<hr/>


### Get Todos - Get
```
http://localhost:8080/getTodos
```

### Delete Todo - Delete
```
http://localhost:8080/deleteTodo/id
```
