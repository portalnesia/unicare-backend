/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package model

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	BaseModel
	NIK      string     `json:"nik" gorm:"column:nik;type:varchar(16);not null;index:idx_customer"`
	Name     string     `json:"name" gorm:"column:name;type:varchar(255);not null;index:idx_customer"`
	Active   bool       `json:"active" gorm:"column:active;type:boolean;default:false;index:idx_customer"`
	NoBPJS   string     `json:"bpjs_number" gorm:"column:bpjs_number;type:varchar(255);nullable;index:idx_customer"`
	Birthday *time.Time `json:"birthday" gorm:"column:birthday;type:date"`

	Roles Roles `json:"roles" gorm:"-"`
}

func (c *Customer) AfterFind(_ *gorm.DB) (err error) {
	c.Roles = Roles_CUSTOMER

	return
}

func (Customer) TableName() string {
	return "customers"
}

func (c Customer) GetID() int {
	return c.ID
}

func (c Customer) GetRole() Roles {
	return Roles_CUSTOMER
}
