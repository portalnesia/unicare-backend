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

func (r *Repo) GetCustomerByID(tx *gorm.DB, id int) (*model.Customer, error) {
	var user model.Customer
	err := tx.Model(user).Table(user.TableName()).Where(" id=?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, r.app.Exc.NotFound("customer", fmt.Sprintf("%d", id), "nik")
		}
		return nil, r.app.Exc.Server(err)
	}
	return &user, err
}

func (r *Repo) GetCustomerByNIK(tx *gorm.DB, nik string) (*model.Customer, error) {
	var user model.Customer
	err := tx.Model(user).Table(user.TableName()).Where("nik=?", nik).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, r.app.Exc.NotFound("customer", fmt.Sprintf("%s", nik), "nik")
		}
		return nil, r.app.Exc.Server(err)
	}
	return &user, err
}

func (r *Repo) ListCustomers(tx *gorm.DB, req *entity.Pagination) (user []model.Customer, err error) {
	var (
		u     model.Customer
		total int64
	)
	db := tx.Model(u).Table(u.TableName())

	if len(req.Q) >= 3 {
		db = db.Where("nik LIKE ? OR LOWER(name) LIKE ? OR bpjs_number LIKE ?", "%"+req.Q+"%", "%"+req.Q+"%", "%"+req.Q+"%")
	}

	if req.Active.Valid {
		db = db.Where("active=?", req.Active.Data)
	}

	db.Count(&total)
	_ = db.Limit(req.PageSize).Offset(req.Start).Find(&user)

	req.Total = total

	return
}
