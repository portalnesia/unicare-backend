/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.portalnesia.com/nullable"
	"gorm.io/gorm"
	"net/url"
	"unicare/internal/config"
	"unicare/internal/entity"
	"unicare/internal/model"
	"unicare/internal/server/usecase"
)

type Handler struct {
	app     *config.App
	usecase *usecase.UseCase
}

func New(app *config.App, usecase *usecase.UseCase) *Handler {
	return &Handler{app, usecase}
}

func (h *Handler) getCtx(c *fiber.Ctx) (*model.Context, error) {
	ctx, ok := c.Locals("ctx").(*model.Context)
	if !ok {
		log.Error().Msg("Failed to get sekala context")
		return nil, h.app.Exc.Server()
	}
	return ctx, nil
}

func (h *Handler) bodyParser(c *fiber.Ctx, data interface{}) error {
	if err := c.BodyParser(data); err != nil {
		return h.app.Exc.InvalidParameter("request", err)
	}
	return nil
}

func (h *Handler) getPaginationQuery(c *fiber.Ctx) (*entity.Pagination, error) {
	var req entity.Pagination
	if err := c.QueryParser(&req); err != nil {
		log.Error().Msg("Failed to parse query request")
		return nil, h.app.Exc.InvalidParameter("request", err)
	}
	req.Init()
	return &req, nil
}

func (h *Handler) newHandler(c *fiber.Ctx, handler func(ctx *model.Context, tx *gorm.DB) error) error {
	ctx, err := h.getCtx(c)
	if err != nil {
		return err
	}

	var errorTrx error
	err = h.app.DB.Session(&gorm.Session{NewDB: true}).Transaction(func(tx *gorm.DB) error {
		if errorTrx = handler(ctx, tx); errorTrx != nil {
			return errorTrx
		}
		return nil
	})
	if err != nil {
		return errorTrx
	}
	return err
}

type CookieConfig struct {
	Name     string
	Value    string
	Expires  carbon.Carbon
	HTTPOnly bool
}

func (h *Handler) SetCookie(c *fiber.Ctx, cookie CookieConfig) {
	domain := "localhost"

	if h.app.Env.IsProd {
		if apiUrl, err := url.Parse(viper.GetString("url")); err == nil {
			domain = "." + apiUrl.Host
		}
	}

	c.Cookie(&fiber.Cookie{
		Name:     cookie.Name,
		Value:    cookie.Value,
		Expires:  cookie.Expires.ToStdTime(),
		Domain:   domain,
		HTTPOnly: cookie.HTTPOnly,
		Secure:   h.app.Env.IsProd,
	})
}

type ResponseData struct {
	Data    *nullable.Type[any] `json:"data,omitempty"`
	Message string              `json:"message"`
}

func (h *Handler) Response(c *fiber.Ctx, data any, message ...any) error {
	httpCode := 200

	msg := "Success"
	if len(message) > 0 {
		for _, m := range message {
			switch v := m.(type) {
			case string:
				msg = v
			case int:
				httpCode = v
			}
		}
	}

	var responseData *nullable.Type[any] = nil

	if data == nil {
		responseData = nullable.NewTypePtr(data, true, false)
	} else {
		switch dt := data.(type) {
		case bool:
			if !dt {
				responseData = nil
			} else {
				var temp any = dt
				responseData = nullable.NewTypePtr(temp, true, true)
			}
		default:
			responseData = nullable.NewTypePtr(data, true, true)
		}
	}

	resp := ResponseData{
		Message: msg,
		Data:    responseData,
	}

	return c.Status(httpCode).JSON(resp)
}
