package logic

import (
	"context"
	"fmt"

	"github.com/polnaya-katuxa/ds-lab-01/internal/models"
	"go.uber.org/zap"
)

type PersonLogic struct {
	repository personRepository

	logger *zap.SugaredLogger
}

func New(repository personRepository, logger *zap.SugaredLogger) *PersonLogic {
	return &PersonLogic{
		repository: repository,
		logger:     logger,
	}
}

func (p *PersonLogic) Create(ctx context.Context, person models.Person) (*models.Person, error) {
	p.logger.Debugw("creating person")

	err := person.Validate()
	if err != nil {
		p.logger.Warnw("invalid person", "error", err)
		return nil, fmt.Errorf("validate person: %w", err)
	}

	created, err := p.repository.Create(ctx, person)
	if err != nil {
		p.logger.Errorw("cannot create person", "error", err)
		return nil, fmt.Errorf("create user in repository: %w", err)
	}

	p.logger.Debugw("created person", "id", created.ID)
	return created, nil
}

func (p *PersonLogic) Get(ctx context.Context, id int) (*models.Person, error) {
	p.logger.Debugw("getting person", "id", id)

	person, err := p.repository.Get(ctx, id)
	if err != nil {
		p.logger.Errorw("cannot get person by id", "error", err, "id", id)
		return nil, fmt.Errorf("get person by id in repository: %w", err)
	}

	p.logger.Debugw("got person", "id", person.ID)
	return person, nil
}

func (p *PersonLogic) Delete(ctx context.Context, id int) error {
	p.logger.Debugw("deleting person", "id", id)

	err := p.repository.Delete(ctx, id)
	if err != nil {
		p.logger.Errorw("cannot delete person", "error", err, "id", id)
		return fmt.Errorf("delete person in repository: %w", err)
	}

	p.logger.Debugw("deleted person", "id", id)
	return nil
}

func (p *PersonLogic) Edit(ctx context.Context, patch models.PersonPatch) (*models.Person, error) {
	p.logger.Debugw("editing person", "id", patch.ID)

	person, err := p.repository.Get(ctx, patch.ID)
	if err != nil {
		p.logger.Errorw("cannot find person", "error", err, "id", patch.ID)
		return nil, fmt.Errorf("find person in repository: %w", err)
	}

	person.Merge(patch)

	err = person.Validate()
	if err != nil {
		p.logger.Warnw("invalid person", "error", err, "id", person.ID)
		return nil, fmt.Errorf("validate person: %w", err)
	}

	err = p.repository.Edit(ctx, *person)
	if err != nil {
		p.logger.Errorw("cannot edit person", "error", err, "id", person.ID)
		return nil, fmt.Errorf("edit person in repository: %w", err)
	}

	p.logger.Debugw("edited person", "id", person.ID)

	return person, nil
}

func (p *PersonLogic) GetAll(ctx context.Context) ([]models.Person, error) {
	p.logger.Debugw("getting all persons")

	persons, err := p.repository.GetAll(ctx)
	if err != nil {
		p.logger.Errorw("cannot get all persons", "error", err)
		return nil, fmt.Errorf("get persons from repository: %w", err)
	}

	p.logger.Debugw("got all persons")
	return persons, nil
}

//go:generate mockery --all --with-expecter --exported --output mocks/

type personRepository interface {
	Create(context.Context, models.Person) (*models.Person, error)
	Get(context.Context, int) (*models.Person, error)
	Delete(context.Context, int) error
	Edit(context.Context, models.Person) error
	GetAll(context.Context) ([]models.Person, error)
}
