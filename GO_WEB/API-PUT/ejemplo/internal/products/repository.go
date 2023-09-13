package products

type Repository interface {
	GetByID(id int) (*Product, error)
	Update(id int) (*Product, error)
	Delete(id int) error
}
type SliceBasedRepository struct {
	productc []*Product
}

func (repository SliceBasedRepository) GetByID(id int) (product *Product, err error) {
	for index := range repository.productc {
		if repository.productc[index].ID == id {
			product = repository.productc[index]
			return
		}
	}
	err = ErrProductNotFound
	return
}
