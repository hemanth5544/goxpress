package services

import (
	"github.com/hemanth5544/goxpress/internal/product/dto"
	"github.com/hemanth5544/goxpress/internal/product/model"
	"github.com/hemanth5544/goxpress/internal/product/repository"
)

type ProductServices struct {
	repo *repository.ProductRepository
}

func NewProductServices(repo *repository.ProductRepository) *ProductServices {
	return &ProductServices{repo: repo}
}

func (s *ProductServices) CreateProduct(productRequest dto.ProductRequest) error {

	return s.repo.CreateProduct(productRequest)

}

func (s *ProductServices) DeleteProductById(id int) error {
	return s.repo.DeleteProductById(id)
}

func (s *ProductServices) GetProductById(id int) (*model.Product, error) {
	return s.repo.GetProductById(id)
}

// ? this is simple see above we retun the model.Product it was type stuct
// ! but see we retuning the *[]model.Product here htat will bing all sticts in array
func (s *ProductServices) GetAllProduct() (*[]model.Product, error) {
	return s.repo.GetAllProduct()
}
