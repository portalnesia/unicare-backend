/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package repo

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"unicare/internal/entity"
	"unicare/internal/model"
)

func (r *Repo) GetUserByID(tx *gorm.DB, id int) (*model.User, error) {
	var user model.User
	err := tx.Model(user).Table(user.TableName()).Where(" id=?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, r.app.Exc.NotFound("user", fmt.Sprintf("%d", id))
		}
		return nil, r.app.Exc.Server(err)
	}
	return &user, err
}

func (r *Repo) GetUserByEmail(tx *gorm.DB, roles model.Roles, email string) (*model.User, error) {
	var user model.User
	err := tx.Model(user).Table(user.TableName()).Where("roles = ? AND email = ?", roles, email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, r.app.Exc.NotFound("user", fmt.Sprintf("%s", email), "email")
		}
		return nil, r.app.Exc.Server(err)
	}
	return &user, err
}

func (r *Repo) ListUsers(tx *gorm.DB, req *entity.Pagination, role model.Roles) (user []model.User, err error) {
	var (
		u     model.User
		total int64
	)
	db := tx.Model(u).Table(u.TableName()).Where("roles = ?", role)

	if len(req.Q) >= 3 {
		db = db.Where("LOWER(email) LIKE ? OR LOWER(name) LIKE ?", "%"+req.Q+"%", "%"+req.Q+"%")
	}

	db.Count(&total)
	_ = db.Limit(req.PageSize).Offset(req.Start).Find(&user)

	req.Total = total

	return
}
