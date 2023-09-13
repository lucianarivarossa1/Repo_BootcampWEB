package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewControllerProducts(db []*Product, lastID int) *ControllerProducts {
	return &ControllerProducts{
		db:     db,
		lastID: lastID,
	}
}

type Product struct {
	ID       int
	Name     string
	Type     string
	Price    float64
	Quantity int
}
type ControllerProducts struct {
	db     []*Product
	lastID int
}
type RequestBody struct {
	Name     string  `json:"name,omitempty"`
	Type     string  `json:"type,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}
type Data struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Type     string  `json:"type,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}
type ResponseBody struct {
	Message string `json:"message"`
	Data    *Data  ` json:"data"`
}

func (c *ControllerProducts) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
func (c *ControllerProducts) Save() gin.HandlerFunc {
	//el gin.Context es para recibir las peticiones del cliente y para responder
	return func(ctx *gin.Context) {
		//REQUEST

		//header
		token := ctx.GetHeader("Authorization")
		if token != "123" {
			code := http.StatusUnauthorized
			body := ResponseBody{Message: "invalid token"}
			ctx.JSON(code, body)
			return
		}

		//body
		var reqBody RequestBody
		err := ctx.ShouldBindJSON(&reqBody)
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseBody{
				Message: "invaid request",
				Data:    nil,
			}
			ctx.JSON(code, body)
			return
		}

		//PROCESS
		//Deserializacion
		pr := &Product{
			Name:     reqBody.Name,
			Type:     reqBody.Type,
			Price:    reqBody.Price,
			Quantity: reqBody.Quantity,
		}
		pr.ID = c.lastID + 1
		//guardar en storage
		c.db = append(c.db, pr)
		c.lastID++

		//RESPONSE
		//Serializacion
		code := http.StatusCreated
		body := ResponseBody{
			Message: "Producto creado",
			Data: &Data{
				ID:       pr.ID,
				Name:     pr.Name,
				Type:     pr.Type,
				Price:    pr.Price,
				Quantity: pr.Quantity,
			},
		}
		ctx.JSON(code, body)

	}
}
