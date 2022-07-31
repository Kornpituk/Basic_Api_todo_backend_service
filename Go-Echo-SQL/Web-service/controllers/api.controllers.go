package controllers

// import (
// 	"myapp/model"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )


// func getAllEmployees(c echo.Context) error {
// 	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	var employees []model.Employee
// 	db.Find(&employees)
// 	return c.JSON(http.StatusOK, employees)
// }

// func getEmployee(c echo.Context) error {
// 	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
	
// 	var employees []model.Employee
// 	db.Find(&employees)
// 	return c.JSON(http.StatusOK, employees)
// }

// func createdEmplyee(c echo.Context) error {
// 	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	employee := new(model.Employee)
// 	if err = c.Bind(employee); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	  }

// 	result := db.Create(&employee)
// 	return c.JSON(http.StatusOK, result)
// }

// func updatedEmplyee(c echo.Context) error {
// 	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	employee := new(model.Employee)
// 	updatedEmplyee := new(model.Employee)
// 	if err = c.Bind(employee); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	  }

// 	db.First(&updatedEmplyee, employee.ID)
// 	updatedEmplyee.Fname = employee.Fname
// 	updatedEmplyee.Lname = employee.Lname
// 	updatedEmplyee.Task = employee.Task
// 	db.Save(updatedEmplyee)

// 	return c.JSON(http.StatusOK, updatedEmplyee)
// }

// func deletedEmployee(c echo.Context) error {
// 	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	id := c.Param("id")
// 	var employee model.Employee
// 	db.First(&employee, id)

// 	db.Delete(&employee)
// 	return c.JSON(http.StatusOK, employee)
// }