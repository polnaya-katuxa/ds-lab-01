package openapi

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/polnaya-katuxa/ds-lab-01/internal/generated/openapi"
	"github.com/polnaya-katuxa/ds-lab-01/internal/models"
)

func toPerson(req openapi.PersonRequest) models.Person {
	return models.Person{
		Name:    req.Name,
		Age:     int(req.Age),
		Address: req.Address,
		Work:    req.Work,
	}
}

func fromPersons(p []models.Person) []openapi.PersonResponse {
	resp := make([]openapi.PersonResponse, 0, len(p))
	for _, v := range p {
		resp = append(resp, fromPerson(v))
	}

	return resp
}

func fromPerson(p models.Person) openapi.PersonResponse {
	return openapi.PersonResponse{
		Id:      int32(p.ID),
		Name:    p.Name,
		Age:     int32(p.Age),
		Address: p.Address,
		Work:    p.Work,
	}
}

func processError(c echo.Context, err error, comment string) error {
	err = fmt.Errorf("%s: %w", comment, err)

	switch {
	case errors.Is(err, models.ErrInvalidData):
		var valErrors validator.ValidationErrors
		if errors.As(err, &valErrors) {
			errorMap := make(map[string]string, len(valErrors))
			for _, v := range valErrors {
				errorMap[v.Field()] = v.Error()
			}

			return c.JSON(http.StatusBadRequest, openapi.ValidationErrorResponse{
				Message: err.Error(),
				Errors:  errorMap,
			})
		}
		return c.JSON(http.StatusBadRequest, openapi.ValidationErrorResponse{
			Message: err.Error(),
			Errors:  map[string]string{},
		})
	case errors.Is(err, models.ErrPersonNotFound):
		return c.JSON(http.StatusNotFound, openapi.ErrorResponse{
			Message: err.Error(),
		})
	default:
		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Message: err.Error(),
		})
	}
}
