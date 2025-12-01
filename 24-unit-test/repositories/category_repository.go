package repositories

import entity "github.com/umarhadimukti/go-learn/24-unit-test/entities"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}