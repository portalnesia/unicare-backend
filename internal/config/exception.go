/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package config

import (
	"fmt"
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
		e.Err = err[0]
		e.Description = err[0].Error()
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

func (Exception) NotFound(model, id string, err ...error) *Error {
	return newError(fiber.StatusNotFound, fmt.Sprintf("%s with id %s not found", model, id), err...)
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
