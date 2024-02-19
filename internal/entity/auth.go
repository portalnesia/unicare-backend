/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package entity

import "unicare/internal/model"

type LoginRequest struct {
	Roles    model.Roles `json:"roles" validate:"required"`
	NIK      string      `json:"nik"`
	Email    string      `json:"email"`
	Password string      `json:"password""`
}
