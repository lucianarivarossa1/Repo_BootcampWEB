package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"` //si pongo - en lugar de json:password no exporta en json
}

func main() {

	user := User{
		Id:       123,
		Name:     "Luciana",
		Password: "contraseña",
	}
	//conviero la estructura a json (devuelve bytes y error)
	userAsJson, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(userAsJson))

	//Gin es framework que esta encima del net/http que sirve para crear servidor, da funcionalidades como routing, middlewords, reduce codigo repetitivo, etc
	//net/http es nativo de GO, se usa router Chi

	//creamos servidor
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": "pong",
		})
	})

	//correr servidor
	router.Run("localhost:8080")
}
