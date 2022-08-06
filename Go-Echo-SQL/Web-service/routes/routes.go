package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
	
)



func Init() *echo.Echo{
	e := echo.New()

	e.GET("/todos", controllers.GetAllTodos)
	e.POST("/todo", controllers.CreatedTodo)
	e.GET("/todo/:id", controllers.GetTodo)
	e.PUT("/todo/:id", controllers.EditedTodo)
	e.DELETE("/todo/:id", controllers.DeletedTodo)

	return e
}