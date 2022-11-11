package routes

import (
	"api-catalog/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()
	route.Static("/image", "./image")
	// route.POST("/todo", controllers.CreateTodo)
	route.GET("/product", controllers.GetAllProduct)
	route.GET("/product/:id_product", controllers.GetProduct)
	route.PUT("/product/:id_product", controllers.UpdateProduct)
	// route.DELETE("/todo/:idTodo", controllers.DeleteTodo)

	route.Run("10.19.23.18:8081")
}
