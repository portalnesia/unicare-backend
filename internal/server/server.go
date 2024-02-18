/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package server

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog/log"
	"unicare/internal/config"
	"unicare/internal/server/handler"
	"unicare/internal/server/middleware"
	"unicare/internal/server/repo"
	"unicare/internal/server/routes"
	"unicare/internal/server/usecase"
)

func New(app *config.App) (fiberApp *fiber.App) {
	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ParserType:        registerDecoder(),
		ZeroEmpty:         true,
	})

	fiberApp = fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024,
		ErrorHandler: func(c *fiber.Ctx, errHandler error) error {
			err := app.Exc.Server()
			var e *config.Error
			if errors.As(errHandler, &e) {
				if !app.Env.IsProd {
					log.Debug().Err(e.Err).Int("status", e.Status).Stack().Msg("Error")
				}

				if e.Message == "Something went wrong" && e.Err != nil {
					fmt.Println(e.Err)
				}
				err = e
				return c.Status(err.Status).JSON(fiber.Map{"data": nil, "error": e})
			} else {
				log.Error().Err(errHandler).Stack().Msg("Erro")
				err.Err = errHandler
				err.Description = errHandler.Error()
				return c.Status(err.Status).JSON(fiber.Map{"data": nil, "error": err})
			}
		},
		AppName: "Unicare by Northbit",
	})

	fiberApp.Use(recover2.New(recover2.Config{
		EnableStackTrace: true,
	}))

	if !app.Env.IsProd {
		// Logger
		fiberApp.Use(logger.New())
	}

	fiberApp.Use(earlydata.New())

	// Compress
	fiberApp.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// Etag
	fiberApp.Use(etag.New())

	// Request ID
	fiberApp.Use(requestid.New())

	r := repo.New(app)
	u := usecase.New(app, r)
	h := handler.New(app, u)

	// middleware
	m := middleware.New(app, r)
	fiberApp.Use(m.Init)

	// Cors
	fiberApp.Use(m.Cors())

	// idempotency
	fiberApp.Use(idempotency.New())

	api := fiberApp.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "API Uptime"})
	})

	routes.New(api, h, m)

	fiberApp.Use(func(c *fiber.Ctx) error {
		err := app.Exc.EndpointNotFound()
		return c.Status(err.Status).JSON(fiber.Map{"data": nil, "error": err})
	})

	return
}
