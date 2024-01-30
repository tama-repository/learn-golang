package service

import (
	"errors"
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service *ProductService) GetById(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("product not found")
	} else {
		return product, nil
	}
}
