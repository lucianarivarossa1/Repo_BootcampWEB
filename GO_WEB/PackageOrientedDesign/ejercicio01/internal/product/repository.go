package product

import (
	"gobases/GO_WEB/PackageOrientedDesign/ejercicio01/internal/domain"
)

type ProductRepository struct {
	productsDB []domain.Product
}

func (r *ProductRepository) GetAllProducts() []domain.Product {
	return r.productsDB
}
func (r *ProductRepository) GetById(id int) domain.Product {
	return r.productsDB[id-1]
}
