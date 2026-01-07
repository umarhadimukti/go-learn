package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	post_mock "github.com/umarhadimukti/go-learn/25-go-mock-test/mock"
)

var mockPostRepo *post_mock.MockPostRepo = new(post_mock.MockPostRepo)
var postService *post_mock.PostService = &post_mock.PostService{Repo: mockPostRepo}

func TestGetPostByTitle_Success(t *testing.T) {
	// create sample data
	post := post_mock.Post{
		ID:      1,
		Title:   "How to get billion's dollar",
		Author:  "test@example.test",
		Content: "lorem",
	}
	// create mock with sample data
	mockPostRepo.On("FindByTitle", "How to get billion's dollar").Return(&post, nil)
	// create test service
	result, err := postService.GetPostByTitle("How to get billion's dollar")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestGetPostByTitle_NotFound(t *testing.T) {
	mockPostRepo.On("FindByTitle", "Lorem Title").Return(nil, assert.AnError)
	result, err := postService.GetPostByTitle("Lorem Title")
	assert.NotNil(t, err)
	assert.Nil(t, result)
}
