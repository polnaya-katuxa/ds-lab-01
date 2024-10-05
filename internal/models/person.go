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

type PersonPatch struct {
	ID      int
	Name    string
	Age     int
	Address string
	Work    *string
}

func (p *Person) Merge(patch PersonPatch) {
	if patch.Name != "" {
		p.Name = patch.Name
	}

	if patch.Address != "" {
		p.Address = patch.Address
	}

	if patch.Age != 0 {
		p.Age = patch.Age
	}

	if patch.Work != nil {
		p.Work = patch.Work
	}
}
