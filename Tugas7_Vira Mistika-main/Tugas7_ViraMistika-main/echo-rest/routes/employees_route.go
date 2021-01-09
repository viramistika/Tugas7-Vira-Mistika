package routes

import (
	"TugasGO7RichyAzura/echo-rest/controllers"

	"github.com/labstack/echo"
)

//CustomersRoute ...
func EmployeesRoute(g *echo.Group) {

	g.POST("/list", controllers.FetchAllEmployees)

	g.POST("/add", controllers.AddEmployees)

	g.POST("/update", controllers.UpdateEmployees)

	g.POST("/delete", controllers.DeleteEmployees)

}
