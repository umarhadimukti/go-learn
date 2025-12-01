package services

import (
	category_repository "github.com/umarhadimukti/go-learn/24-unit-test/repositories"
	entity "github.com/umarhadimukti/go-learn/24-unit-test/entities"
	"errors"
)

type CategoryService struct {
	Repository category_repository.CategoryRepository
}

func (categoryService CategoryService) GetDetail(id string) (*entity.Category, error) {
	category := categoryService.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("no category found")
	}
	return category, nil
}