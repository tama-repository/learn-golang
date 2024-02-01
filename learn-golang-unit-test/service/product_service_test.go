package service

import (
	"learn-golang-unit-test/entity"
	"learn-golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetById(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)
	product, err := productService.GetById("1")
	assert.Nil(t, product)
	assert.NotNil(t, err)
}

func TestProductService(t *testing.T) {
	product := entity.Product{
		Id: "2",
		Name: "Product-1",
		Price: "20000",
		Description: "This is Product-1",
	}

	productRepository.Mock.On("FindById", "2").Return(product)
	p, err := productService.GetById("2")
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, product.Id, p.Id)
	assert.Equal(t, product.Name, p.Name)
	assert.Equal(t, product.Price, p.Price)
	assert.Equal(t, product.Description, p.Description)
}
