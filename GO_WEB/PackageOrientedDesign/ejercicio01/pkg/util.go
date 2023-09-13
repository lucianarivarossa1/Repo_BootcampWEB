package pkg

import (
	"encoding/json"
	"gobases/GO_WEB/PackageOrientedDesign/ejercicio01/internal/domain"
	"io"
	"log"
	"os"
)

// cargar datos de json a bbd
func FillfilProductosDB(path string, slice *[]domain.Product) {
	data, err := os.Open("products.json")
	if err != nil {
		log.Fatal(err)
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(dataRead, &slice)
}

func MarshalingData(object any) (data []byte) {
	return
}
