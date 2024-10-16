package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada de Usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Camada de Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/product/:id", ProductController.GetProductById)
	server.POST("/create", ProductController.CreateProduct)
	server.PUT("/update", ProductController.UpdateProduct)
	server.DELETE("/delete/:id", ProductController.DeleteProductById)

	server.Run(":8000")
}
