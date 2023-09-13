package product

import (
	"gobases/GO_WEB/PackageOrientedDesign/ejercicio01/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// hacer una esctrucuta que represente a mi setvice que tenaga una istancia a mi resop que necesito consultar
type ProductService struct {
	repository ProductRepository
}

func Ping(ctx *gin.Context) {
	ctx.String(200, "Pong")
}
func (s *ProductService) GetAllProducts(ctx *gin.Context) {
	data := pkg.MarshalingData(s.repository.GetAllProducts())
	ctx.Data(http.StatusOK, "application/json", data)
}
