package models

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidData    = errors.New("invalid data")
	ErrPersonNotFound = errors.New("person not found")
)

type Person struct {
	ID      int     `validate:"omitempty,gte=0"`
	Name    string  `validate:"required"`
	Age     int     `validate:"required,gte=0"`
	Address string  `validate:"required"`
	Work    *string `validate:"omitempty,gt=0"`
}

func (p *Person) Validate() error {
	err := validator.New().Struct(p)
	if err != nil {
		return fmt.Errorf("validate person: %w (%w)", err, ErrInvalidData)
	}

	return nil
}
