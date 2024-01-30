package service

import (
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
