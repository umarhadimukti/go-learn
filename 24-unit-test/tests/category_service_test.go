package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	repository "github.com/umarhadimukti/go-learn/24-unit-test/repositories"
	service "github.com/umarhadimukti/go-learn/24-unit-test/services"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = service.CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	// mock program
	categoryRepository.Mock.On("FindById", "2").Return(nil)

	category, error := categoryService.GetDetail("2")
	assert.Nil(t, category)
	assert.NotNil(t, error)
}

