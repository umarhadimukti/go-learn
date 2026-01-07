package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/umarhadimukti/go-learn/24-unit-test/entities"
	repository "github.com/umarhadimukti/go-learn/24-unit-test/repositories"
	service "github.com/umarhadimukti/go-learn/24-unit-test/services"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = service.CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// mock program
	categoryRepository.Mock.On("FindById", "2").Return(nil)

	category, error := categoryService.GetDetail("2")
	assert.Nil(t, category)
	assert.NotNil(t, error)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	categoryMock := entities.Category{
		Id: "1",
		Name: "Sport",
	}
	// mock program
	categoryRepository.Mock.On("FindById", "1").Return(categoryMock)

	result, err := categoryService.GetDetail("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

