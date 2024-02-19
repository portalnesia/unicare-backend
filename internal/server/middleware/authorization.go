/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.portalnesia.com/nullable"
	"unicare/internal/model"
	"unicare/pkg/constant"
	"unicare/pkg/jwt"
)

func (m *Middleware) Authorization(c *fiber.Ctx) error {
	ctx := c.Locals("ctx").(*model.Context)
	auth := c.Cookies(constant.Cookie_AUTH)
	if auth == "" {
		return m.app.Exc.Unauthorized()
	}

	verified, data, _ := jwt.Verify(auth)
	if !verified {
		return m.app.Exc.Unauthorized()
	}

	if rl, ok := data["rl"].(string); ok {
		if rl == "PATIENT" {
			return m.authorizationCustomer(c)
		}
	} else {
		return m.app.Exc.Unauthorized()
	}

	id := data["sub"].(float64)
	if u, err := m.repo.GetUserByID(m.app.DB, int(id)); err == nil && u != nil {
		ctx.User = nullable.NewType(*u, true, true)
		c.Locals("ctx", ctx)
		return c.Next()
	}

	return m.app.Exc.Unauthorized()
}

func (m *Middleware) authorizationCustomer(c *fiber.Ctx) error {
	ctx := c.Locals("ctx").(*model.Context)
	auth := c.Cookies(constant.Cookie_AUTH)
	if auth == "" {
		return m.app.Exc.Unauthorized()
	}

	verified, data, _ := jwt.Verify(auth)
	if !verified {
		return m.app.Exc.Unauthorized()
	}

	id := data["sub"].(float64)
	if u, err := m.repo.GetCustomerByID(m.app.DB, int(id)); err == nil && u != nil {
		ctx.Customer = nullable.NewType(*u, true, true)
		c.Locals("ctx", ctx)
		return c.Next()
	}

	return m.app.Exc.Unauthorized()
}

func (m *Middleware) OnlyLogin(role ...model.Roles) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Locals("ctx").(*model.Context)

		if !ctx.User.Valid {
			if !ctx.Customer.Valid {
				return m.app.Exc.Unauthorized()
			}
		}

		if len(role) > 0 {
			found := false
			for _, r := range role {
				if r == model.Roles_CUSTOMER && ctx.Customer.Valid {
					found = true
					break
				}

				if ctx.User.Data.Roles == r {
					found = true
					break
				}
			}
			if !found {
				return m.app.Exc.Forbidden()
			}
		}

		return c.Next()
	}
}
