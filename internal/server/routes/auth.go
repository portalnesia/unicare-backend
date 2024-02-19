/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package routes

import "github.com/gofiber/fiber/v2"

func (r *Routes) auth(app fiber.Router) {
	route := app.Group("/auth")

	route.Post("/login", r.handler.Login)
}
