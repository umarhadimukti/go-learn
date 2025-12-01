package repositories

import (
	"github.com/stretchr/testify/mock"
	category_entity "github.com/umarhadimukti/go-learn/24-unit-test/entities"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *category_entity.Category {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	category := args.Get(0).(category_entity.Category)
	return &category
}