package controllers

import (
	"myapp/model"
	"net/http"

	"myapp/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"


	"strconv"
)

// function Data Bases Connect SQL
// func init(){
// 	database.ConnectDB()
// }



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
	return c.JSON(http.StatusOK, "Not Found ID "+c.Param("id"))
}
	

// Edit Todo 
func EditedTodo (c echo.Context) error {

	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}


	id, _ := strconv.Atoi(c.Param("id")) 
	Todo := new(model.Todo)

	updatedTodo := new(model.Todo)
	if err = c.Bind(Todo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

	db.First(&updatedTodo, id)
	
	if int(updatedTodo.ID) == id{
		updatedTodo.First_name = Todo.First_name
		updatedTodo.Last_name = Todo.Last_name
		updatedTodo.Task = Todo.Task
		db.Save(updatedTodo)
		return c.JSON(http.StatusOK, updatedTodo)
	}

		
	return c.JSON(http.StatusOK, "Edite Todo Fail!" + "  Not Found ID "+c.Param("id"))
	

}


//Delete Todo 
func DeletedTodo (c echo.Context) error {

	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var todo model.Todo
	database.DB.First(&todo, id)

	if int(todo.ID) == id{
		db.Where("ID=?",id).Delete(&todo)
		return c.JSON(http.StatusOK, todo)
	}
	return c.JSON(http.StatusOK, "Delete Todo Fail!" + "  Not Found ID "+c.Param("id") )
}
