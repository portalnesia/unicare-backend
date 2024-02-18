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
	"gorm.io/gorm"
	"unicare/internal/entity"
	"unicare/internal/model"
	"unicare/pkg/constant"
	"unicare/pkg/jwt"
)

func (h *Handler) Login(c *fiber.Ctx) error {
	return h.newHandler(c, func(ctx *model.Context, tx *gorm.DB) error {
		var req entity.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return h.app.Exc.InvalidParameter("request", err)
		}
		resp, err := h.usecase.Login(tx, req)
		if err != nil {
			return err
		}

		claims := map[string]any{
			"rl": model.RolesToString[resp.GetRole()],
		}
		token, err := jwt.GenerateJWT(resp.GetID(), claims)
		if err != nil {
			return h.app.Exc.Server(err)
		}
		h.SetCookie(c, CookieConfig{
			Name:    constant.Cookie_AUTH,
			Value:   token,
			Expires: carbon.Now("Asia/Jakarta").AddHours(6),
		})

		return h.Response(c, resp)
	})
}
