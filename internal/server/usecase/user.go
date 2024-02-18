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

func (u *UseCase) ListUsers(tx *gorm.DB, ctx *model.Context, req *entity.Pagination, role model.Roles) (any, error) {
	if err := u.app.Validator.Struct(req); err != nil {
		return nil, u.app.Exc.Validator(err)
	}

	if role == model.Roles_CUSTOMER {
		users, err := u.repo.ListCustomers(tx, req)
		if err != nil {
			return nil, err
		}

		return users, nil
	} else {
		users, err := u.repo.ListUsers(tx, req, role)
		if err != nil {
			return nil, err
		}

		return users, nil
	}
}
