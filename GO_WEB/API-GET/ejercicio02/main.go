package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre   string `json:"nombre,omitempty"`
	Apellido string `json:"apellido,omitempty"`
}
type ResponseSaludo struct {
	Saludo   string
	Nombre   string
	Apellido string
}

func SaludarUno(saludo string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := saludo
		ctx.JSON(http.StatusOK, body)
	}
}
func SaludarDos(nombre, apellido string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := ResponseSaludo{
			Saludo:   "hola",
			Nombre:   nombre,
			Apellido: apellido,
		}
		ctx.JSON(http.StatusOK, body)
	}
}
func main() {
	p := Persona{
		"Luciana",
		"Rivarossa",
	}
	jsonPersona, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	saludo := "Hola " + string(jsonPersona)

	router := gin.Default()
	router.GET("/saludoUno", SaludarUno(saludo))
	router.GET("/saludoDos", SaludarDos(p.Nombre, p.Apellido))
	router.Run("localhost:8080")
}
