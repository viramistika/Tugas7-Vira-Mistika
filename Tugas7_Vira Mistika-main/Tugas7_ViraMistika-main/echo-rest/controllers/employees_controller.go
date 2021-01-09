package controllers

import (
	"Tugas7_ViraMistika/echo-rest/models"
	"net/http"

	"github.com/labstack/echo"
)

//FetchAllCustomers ...
func FetchAllEmployees(c echo.Context) (err error) {

	result, err := models.FetchEmployees()

	return c.JSON(http.StatusOK, result)
}

//AddEmployees
func AddEmployees(c echo.Context) (err error) {

	result, err := models.AddEmployee(c)

	return c.JSON(http.StatusOK, result)
}

//UpdateEmployees
func UpdateEmployees(c echo.Context) (err error) {

	result, err := models.UpdateEmployee(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteEmployees ...
func DeleteEmployees(c echo.Context) (err error) {

	result, err := models.DeleteEmployee(c)

	return c.JSON(http.StatusOK, result)
}
