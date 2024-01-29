package main

import (
	"github.com/RyanRamadhan11/Shop_RestApi_App/controllers/productController"
	"github.com/RyanRamadhan11/Shop_RestApi_App/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productController.Index)
	r.GET("/api/product/:id", productController.Show)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product/:id", productController.Update)
	r.DELETE("/api/product", productController.Delete)

	r.Run()
}