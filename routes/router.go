package routes

import (
	"github.com/gin-gonic/gin"
	"wallester/controllers"
)

func Attach(router *gin.Engine,controller controllers.CustomerInterface) {
	// Views
	router.GET("/", controller.TableView)
	router.POST("/", controller.TableView)
	router.GET("/edit/:id", controller.EditView)
	router.GET("/add", controller.CreateView)
	router.GET("/search", controller.SearchView)
	//Api
	routesApi := router.Group("/api")
	{
		routesApi.POST("/customers/", controller.ShowCustomer)
		routesApi.PUT("/customers/:id", controller.UpdateCustomer)
		routesApi.POST("/customers/add", controller.InsertCustomer)
		routesApi.GET("/customers/search", controller.SearchCustomer)
	}
}
