/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package config

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Status      int    `json:"status"` // http status
	Message     string `json:"message"`
	Description string `json:"description"`
	Err         error  `json:"reason"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s (%s)", e.Message, e.Description)
}

func newError(status int, message string, err ...error) *Error {
	e := &Error{
		Status:  status,
		Message: message,
	}
	if len(err) > 0 {
		if err[0] != nil {
			e.Err = err[0]
			e.Description = err[0].Error()
		}
	}
	return e
}

type Exception struct{}

func (Exception) Server(err ...error) *Error {
	return newError(fiber.StatusServiceUnavailable, "Something went wrong", err...)
}
func (Exception) EndpointNotFound() *Error {
	return newError(fiber.StatusNotFound, "Endpoint not found")
}

func (Exception) NotFound(model, id string, data ...interface{}) *Error {
	idMsg := "id"
	var err error = nil
	if len(data) > 0 {
		for _, d := range data {
			switch v := d.(type) {
			case string:
				idMsg = v
			case error:
				err = v
			}
		}
	}

	return newError(fiber.StatusNotFound, fmt.Sprintf("%s with %s %s not found", model, idMsg, id), err)
}

func (Exception) BadParameter(params string, err ...error) *Error {
	return newError(fiber.StatusBadRequest, fmt.Sprintf("Missing %s parameter", params), err...)
}

func (Exception) InvalidParameter(params string, err ...error) *Error {
	return newError(fiber.StatusBadRequest, fmt.Sprintf("Invalid %s parameter", params), err...)
}

func (Exception) Unauthorized(err ...error) *Error {
	return newError(fiber.StatusUnauthorized, "Unauthorized", err...)
}

func (Exception) Forbidden(err ...error) *Error {
	return newError(fiber.StatusForbidden, "You dont have access to perform this action", err...)
}

func (e Exception) Validator(err error) *Error {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		er := fmt.Errorf("field validation error on the '%s' tag", errs[0].Tag())
		return e.InvalidParameter(errs[0].Field(), er)
	}

	return newError(fiber.StatusBadRequest, "Invalid request", err)
}
