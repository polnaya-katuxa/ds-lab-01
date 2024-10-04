package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/polnaya-katuxa/ds-lab-01/internal/models"
	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PersonRepository {
	return &PersonRepository{
		db: db,
	}
}

func (p *PersonRepository) Create(ctx context.Context, person models.Person) (*models.Person, error) {
	err := p.db.Table("persons").Create(&person).Error
	if err != nil {
		return nil, fmt.Errorf("create person in db: %w", err)
	}

	return &person, nil
}

func (p *PersonRepository) Get(ctx context.Context, id int) (*models.Person, error) {
	var person models.Person
	err := p.db.Table("persons").Where("id = ?", id).First(&person).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("get person from db: %w", models.ErrPersonNotFound)
		}

		return nil, fmt.Errorf("get person from db: %w", err)
	}

	return &person, nil
}

func (p *PersonRepository) Delete(ctx context.Context, id int) error {
	err := p.db.Table("persons").Where("id = ?", id).Delete(&models.Person{}, id).Error
	if err != nil {
		return fmt.Errorf("delete person from db: %w", err)
	}

	return nil
}

func (p *PersonRepository) Edit(ctx context.Context, person models.Person) error {
	res := p.db.Table("persons").Where("id = ?", person.ID).Updates(map[string]any{
		"name":    person.Name,
		"age":     person.Age,
		"address": person.Address,
		"work":    person.Work,
	})
	if res.Error != nil {
		return fmt.Errorf("update person in db: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("update person in db: %w", models.ErrPersonNotFound)
	}

	return nil
}

func (p *PersonRepository) GetAll(ctx context.Context) ([]models.Person, error) {
	var persons []models.Person
	err := p.db.Table("persons").Find(&persons).Error
	if err != nil {
		return nil, fmt.Errorf("get persons from db: %w", err)
	}

	return persons, nil
}
