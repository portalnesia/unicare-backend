/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
	"unicare/internal/model"
)

func (h *Handler) GetMe(c *fiber.Ctx) error {
	return h.newHandler(c, func(ctx *model.Context, tx *gorm.DB) error {
		if ctx.User.Valid {
			return h.Response(c, ctx.User)
		} else if ctx.Customer.Valid {
			return h.Response(c, ctx.Customer)
		}
		return h.Response(c, nil)
	})
}

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	return h.newHandler(c, func(ctx *model.Context, tx *gorm.DB) error {
		role := c.Params("role")
		req, err := h.getPaginationQuery(c)
		if err != nil {
			return err
		}

		r, ok := model.StringToRoles[strings.ToUpper(role)]
		if !ok {
			return h.app.Exc.InvalidParameter("role")
		}

		data, _ := h.usecase.ListUsers(tx, ctx, req, r)
		resp := req.NewPagination(data)

		return h.Response(c, resp)
	})
}
