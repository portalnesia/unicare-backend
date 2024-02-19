/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"unicare/internal/config"
	"unicare/internal/model"
	"unicare/internal/server/repo"
)

type Middleware struct {
	app  *config.App
	repo *repo.Repo
}

func New(app *config.App, repo *repo.Repo) *Middleware {
	return &Middleware{app, repo}
}

func (m *Middleware) Init(c *fiber.Ctx) error {
	c.Set("Vary", "Accept-Encoding")

	ip := setUpIP(m.app, c)

	context := model.Context{
		IP:     ip,
		Method: c.Method(),
	}
	c.Locals("ctx", &context)
	return c.Next()
}

func setUpIP(app *config.App, c *fiber.Ctx) string {
	ip := ""
	if app.Env.IsProd {
		ip = c.Get("cf-connecting-ip", "")
		if ip == "" {
			ip = c.IP()
		}
	} else {
		ip = c.IP()
	}
	return ip
}
