/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"net/url"
	"strings"
)

func (m *Middleware) Cors() fiber.Handler {
	allowed := func(origin string) bool {
		if !m.app.Env.IsProd {
			return true
		}

		link, err := url.Parse(origin)
		if err != nil {
			return false
		}

		domain, err := url.Parse(viper.GetString("url"))
		if err != nil {
			return false
		}

		return strings.Contains(link.Host, domain.Host)
	}
	return cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginsFunc: allowed,
	})
}
