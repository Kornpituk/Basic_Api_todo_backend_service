package controllers

import (
	"myapp/model"
	"net/http"


	"github.com/labstack/echo/v4"
	"myapp/database"

	"strconv"
)

// function Data Bases Connect SQL
func init(){
	database.ConnectDB()
}



//Get All Todos
func GetAllTodos(c echo.Context) error {
	var todo []model.Todo
	database.DB.Find(&todo)
	return c.JSON(http.StatusOK, todo)
}


//Create Todo 
func CreatedTodo (c echo.Context) error {
	todo := new(model.Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	database.DB.Create(&todo)
	return c.JSON(http.StatusOK, todo)
  }



//Get Todo with ID
func GetTodo (c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id")) 
	var todo model.Todo

	database.DB.First(&todo, id)
	if int(todo.ID) == id{
		return c.JSON(http.StatusOK, todo)
	}
	return c.JSON(http.StatusOK, "Not Found ID Todo!")
}
	

// Edit Todo 
func EditedTodo (c echo.Context) error {
	var err error
	id, _ := strconv.Atoi(c.Param("id")) 
	Todo := new(model.Todo)

	updatedTodo := new(model.Todo)
	if err = c.Bind(Todo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

	database.DB.First(&updatedTodo, id)
	if int(updatedTodo.ID) == id{
		updatedTodo.Fname = Todo.Fname
		updatedTodo.Lname = Todo.Lname
		updatedTodo.Task = Todo.Task
		database.DB.Save(updatedTodo)
		return c.JSON(http.StatusOK, updatedTodo)
	}

		
	return c.JSON(http.StatusOK, "Edite Todo Fail!")

}


//Delete Todo 
func DeletedTodo (c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	var todo model.Todo
	database.DB.First(&todo, id)

	if int(todo.ID) == id{
		database.DB.Where("ID=?",id).Delete(&todo)
		return c.JSON(http.StatusOK, todo)
	}
	return c.JSON(http.StatusOK, "Delete Todo Fail!")
}
