package main

import "github.com/gin-gonic/gin"

func SayHello(ctx *gin.Context) {
	ctx.String(200, "Hello World")
}
func SayHelloWithName(ctx *gin.Context) {
	ctx.String(200, "Hello "+ctx.Param("nombre"))
}
func SayHelloWithQueryParams(ctx *gin.Context) {
	ctx.String(200, "Hello this query params is: "+ctx.Query("age"))
}

func main() {
	server := gin.Default()
	helloPaths := server.Group("/hello")
	helloPaths.GET("", SayHello)
	helloPaths.GET("/:nombre", SayHelloWithName)
	helloPaths.GET("/queryParams", SayHelloWithQueryParams)
	//localhost: 8084
	server.Run(":8084")

}
