/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package routes

import (
	"github.com/gofiber/fiber/v2"
	"unicare/internal/model"
)

func (r *Routes) users(app fiber.Router) {
	route := app.Group("/user", r.middle.Authorization)

	route.Get("/me", r.middle.OnlyLogin(), r.handler.GetMe)
	route.Get("/:role", r.middle.OnlyLogin(model.Roles_ADMIN), r.handler.ListUsers)

}
