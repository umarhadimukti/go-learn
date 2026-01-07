package mock

import (
	"github.com/stretchr/testify/mock"
)

// First, create Post struct
type Post struct {
	ID int
	Title, Author, Content string
}
// Second, create interface PostRepository that will be used to register the method
type PostRepository interface {
	FindByTitle(title string) (*Post, error)
}
// Third, create service to be tested
type PostService struct {
	Repo PostRepository
}
func (p *PostService) GetPostByTitle(title string) (*Post, error) {
	post, err := p.Repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return post, nil
}
// Fourth, create mock with testify/mock
type MockPostRepo struct {
	mock.Mock
}
func (m *MockPostRepo) FindByTitle(title string) (*Post, error) {
	args := m.Called(title)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	post := args.Get(0).(*Post)
	return post, nil
}