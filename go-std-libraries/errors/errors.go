package errors

import (
	"errors"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetStudentById(id int) (string, error) {
	if id != 1 {
		return "", NotFoundError
	}
	if id <= 0 {
		return "", ValidationError
	}
	return "Student Name", nil
}