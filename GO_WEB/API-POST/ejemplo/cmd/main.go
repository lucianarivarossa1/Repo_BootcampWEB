package main

import (
	"gobases/GO_WEB/API-POST/ejemplo/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	//env
	//dependencias
	//Creo el slice
	db := make([]*handlers.Product, 0)

	//nueva instancia al controller
	ct := handlers.NewControllerProducts(db, 0)

	//nuevo server
	rt := gin.New()

	//middlewares
	rt.Use(gin.Logger())
	rt.Use(gin.Recovery())

	//routes POST
	//productos
	rt.POST("/products", ct.Save())

	//run
	if err := rt.Run(":8080"); err != nil {
		panic(err)
	}
}
