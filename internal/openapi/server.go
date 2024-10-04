package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/polnaya-katuxa/ds-lab-01/internal/generated/openapi"
	"github.com/polnaya-katuxa/ds-lab-01/internal/models"
)

type Server struct {
	personLogic personLogic
}

func New(personLogic personLogic) *Server {
	return &Server{
		personLogic: personLogic,
	}
}

func (s *Server) ListPersons(c echo.Context) error {
	persons, err := s.personLogic.GetAll(c.Request().Context())
	if err != nil {
		return processError(c, err, "cannot get persons list")
	}

	return c.JSON(http.StatusOK, fromPersons(persons))
}

func (s *Server) CreatePerson(c echo.Context) error {
	var req openapi.PersonRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return processError(c, err, "cannot unmarshal request body")
	}

	person := toPerson(req)
	created, err := s.personLogic.Create(c.Request().Context(), person)
	if err != nil {
		return processError(c, err, "cannot create user")
	}

	c.Response().Header().Add("Location", fmt.Sprintf("/api/v1/persons/%d", created.ID))
	return c.NoContent(http.StatusCreated)
}

func (s *Server) DeletePerson(c echo.Context, id int32) error {
	err := s.personLogic.Delete(c.Request().Context(), int(id))
	if err != nil {
		return processError(c, err, "cannot delete user")
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (s *Server) GetPerson(c echo.Context, id int32) error {
	person, err := s.personLogic.Get(c.Request().Context(), int(id))
	if err != nil {
		return processError(c, err, "cannot get user by id")
	}

	return c.JSON(http.StatusOK, fromPerson(*person))
}

func (s *Server) EditPerson(c echo.Context, id int32) error {
	var req openapi.PersonRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return processError(c, err, "cannot unmarshal request body")
	}

	person := toPerson(req)
	person.ID = int(id)

	updated, err := s.personLogic.Edit(c.Request().Context(), person)
	if err != nil {
		return processError(c, err, "cannot edit user")
	}

	return c.JSON(http.StatusOK, fromPerson(*updated))
}

type personLogic interface {
	Create(context.Context, models.Person) (*models.Person, error)
	Get(context.Context, int) (*models.Person, error)
	Delete(context.Context, int) error
	Edit(context.Context, models.Person) (*models.Person, error)
	GetAll(context.Context) ([]models.Person, error)
}
