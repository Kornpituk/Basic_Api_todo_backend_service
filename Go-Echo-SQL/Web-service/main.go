package main

import (
	"myapp/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// "myapp/controllers"

)


func main() {
	//Conect MySQL DATABASE
	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	e := echo.New()
	// Get All Todos
	e.GET("/todos", func(c echo.Context) error {
		var todo []model.Todo
		db.Find(&todo)
		return c.JSON(http.StatusOK, todo)
	})
	
	//Get Todo with ID
	e.GET("/todo/:id", func(c echo.Context) error {
		id := c.Param("id")
		var employee model.Todo
		db.First(&employee, id)
	  return c.JSON(http.StatusOK, employee)
	})

	//Create  Todo
	e.POST("/todo", func(c echo.Context) error {
		employee := new(model.Todo)
		if err = c.Bind(employee); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		  }

		result := db.Create(&employee)
		return c.JSON(http.StatusOK, result)
	  })
	
	//Delete Todo 
	e.DELETE("/todo/:id", func(c echo.Context) error {
		id := c.Param("id")
		var todo model.Todo
		db.First(&todo, id)

		db.Delete(&todo)
		return c.JSON(http.StatusOK, todo)
	})

	// Edit Todo 
	e.PUT("/todo/:id", func(c echo.Context) error {
		id := c.Param("id")
		Todo := new(model.Todo)
		updatedTodo := new(model.Todo)
		if err = c.Bind(Todo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		  }

		db.First(&updatedTodo, id)
		updatedTodo.Fname = Todo.Fname
		updatedTodo.Lname = Todo.Lname
		updatedTodo.Task = Todo.Task
		db.Save(updatedTodo)

		return c.JSON(http.StatusOK, updatedTodo)
	  })



	e.Logger.Fatal(e.Start(":1323"))
}

