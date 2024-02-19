/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package usecase

import (
	"gorm.io/gorm"
	"unicare/internal/entity"
	"unicare/internal/model"
)

func (u *UseCase) Login(tx *gorm.DB, req entity.LoginRequest) (model.Session, error) {
	if err := u.app.Validator.Struct(req); err != nil {
		return nil, u.app.Exc.Validator(err)
	}

	if req.Roles == model.Roles_CUSTOMER {
		user, err := u.repo.GetCustomerByNIK(tx, req.NIK)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		user, err := u.repo.GetUserByEmail(tx, req.Roles, req.Email)
		if err != nil {
			return nil, err
		}

		if !user.CheckPassword(req.Password) {
			return nil, u.app.Exc.InvalidParameter("password")
		}

		return user, nil
	}
}
