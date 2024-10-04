package logic

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/polnaya-katuxa/ds-lab-01/internal/logic/mocks"
	"github.com/polnaya-katuxa/ds-lab-01/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestPersonLogic_Create(t *testing.T) {
	t.Run("created person", func(t *testing.T) {
		ctx := context.Background()

		personToCreate := models.Person{
			Name:    "Vasya",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		want := &models.Person{
			ID:      1,
			Name:    "Vasya",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Create(ctx, personToCreate).Return(want, nil)

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Create(ctx, personToCreate)
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("invalid person", func(t *testing.T) {
		ctx := context.Background()

		personToCreate := models.Person{
			Name:    "",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		repository := mocks.NewPersonRepository(t)

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Create(ctx, personToCreate)
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("repository error", func(t *testing.T) {
		ctx := context.Background()

		personToCreate := models.Person{
			Name:    "Vasya",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Create(ctx, personToCreate).Return(nil, errors.New("error"))

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Create(ctx, personToCreate)
		require.Error(t, err)
		require.Nil(t, got)
	})
}

func TestPersonLogic_Get(t *testing.T) {
	t.Run("got person", func(t *testing.T) {
		ctx := context.Background()

		id := 1
		want := &models.Person{
			ID:      id,
			Name:    "Vasya",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Get(ctx, id).Return(want, nil)

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Get(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("repository error", func(t *testing.T) {
		ctx := context.Background()

		id := 1
		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Get(ctx, id).Return(nil, errors.New("error"))

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Get(ctx, id)
		require.Error(t, err)
		require.Nil(t, got)
	})
}

func TestPersonLogic_Delete(t *testing.T) {
	t.Run("deleted person", func(t *testing.T) {
		ctx := context.Background()

		id := 1
		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Delete(ctx, id).Return(nil)

		p := New(repository, zap.NewNop().Sugar())
		err := p.Delete(ctx, id)
		require.NoError(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		ctx := context.Background()

		id := 1
		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Delete(ctx, id).Return(errors.New("error"))

		p := New(repository, zap.NewNop().Sugar())
		err := p.Delete(ctx, id)
		require.Error(t, err)
	})
}

func TestPersonLogic_Edit(t *testing.T) {
	t.Run("edited person", func(t *testing.T) {
		ctx := context.Background()

		personToEdit := models.Person{
			ID:      1,
			Name:    "Vasya",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		want := &models.Person{
			ID:      1,
			Name:    "Vasya",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Edit(ctx, personToEdit).Return(nil)

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Edit(ctx, personToEdit)
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("invalid person", func(t *testing.T) {
		ctx := context.Background()

		personToEdit := models.Person{
			ID:      1,
			Name:    "",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		repository := mocks.NewPersonRepository(t)

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Edit(ctx, personToEdit)
		require.Error(t, err)
		require.Nil(t, got)
	})

	t.Run("repository error", func(t *testing.T) {
		ctx := context.Background()

		personToEdit := models.Person{
			ID:      1,
			Name:    "Vasya",
			Age:     18,
			Address: "Moscow",
			Work:    ptr.String("BMSTU"),
		}

		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().Edit(ctx, personToEdit).Return(errors.New("error"))

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.Edit(ctx, personToEdit)
		require.Error(t, err)
		require.Nil(t, got)
	})
}

func TestPersonLogic_GetAll(t *testing.T) {
	t.Run("got person list", func(t *testing.T) {
		ctx := context.Background()

		want := []models.Person{
			{
				ID:      1,
				Name:    "Vasya",
				Age:     18,
				Address: "Moscow",
				Work:    ptr.String("BMSTU"),
			},
			{
				ID:      2,
				Name:    "Masha",
				Age:     3,
				Address: "SPB",
				Work:    ptr.String("Avito"),
			},
		}

		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().GetAll(ctx).Return(want, nil)

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.GetAll(ctx)
		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("repository error", func(t *testing.T) {
		ctx := context.Background()

		repository := mocks.NewPersonRepository(t)
		repository.EXPECT().GetAll(ctx).Return(nil, errors.New("error"))

		p := New(repository, zap.NewNop().Sugar())
		got, err := p.GetAll(ctx)
		require.Error(t, err)
		require.Nil(t, got)
	})
}
