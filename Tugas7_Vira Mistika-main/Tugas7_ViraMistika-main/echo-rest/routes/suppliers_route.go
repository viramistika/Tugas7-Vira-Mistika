package routes

import (
	"TugasGO7RichyAzura/echo-rest/controllers"

	"github.com/labstack/echo"
)

//SuppliersRoute ...
func SuppliersRoute(g *echo.Group) {

	g.POST("/list", controllers.FetchAllSuppliers)

	g.POST("/add", controllers.StoreSuppliers)

	g.POST("/update", controllers.UpdateSuppliers)

	g.POST("/delete", controllers.DeleteSuppliers)

}
