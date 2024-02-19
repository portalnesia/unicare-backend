/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package routes

import (
	"github.com/gofiber/fiber/v2"
	"unicare/internal/server/handler"
	"unicare/internal/server/middleware"
)

type Routes struct {
	handler *handler.Handler
	middle  *middleware.Middleware
}

func New(fiberApp fiber.Router, handler2 *handler.Handler, m *middleware.Middleware) {
	r := &Routes{handler2, m}

	r.auth(fiberApp)
	r.users(fiberApp)
}
