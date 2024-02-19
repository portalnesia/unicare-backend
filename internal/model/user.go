/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package model

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

type Roles int

const (
	Roles_MASTER Roles = iota + 1
	Roles_ADMIN
	Roles_PROVIDER
	Roles_CUSTOMER
)

var (
	RolesToString = map[Roles]string{
		Roles_MASTER:   "MASTER",
		Roles_ADMIN:    "ADMIN",
		Roles_PROVIDER: "PROVIDER",
		Roles_CUSTOMER: "CUSTOMER",
	}
	StringToRoles = map[string]Roles{
		"MASTER":   Roles_MASTER,
		"ADMIN":    Roles_ADMIN,
		"PROVIDER": Roles_PROVIDER,
		"CUSTOMER": Roles_CUSTOMER,
	}
)

// MarshalJSON implements json.Marshaler interface.
func (r Roles) MarshalJSON() ([]byte, error) {
	d := RolesToString[r]
	return json.Marshal(d)
}

// UnmarshalJSON implements json.Marshaler interface.
func (r *Roles) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	// string
	if data[0] == '"' {
		data = data[1 : len(data)-1]
		if d, ok := StringToRoles[string(data)]; ok {
			*r = d
			return nil
		}
		return errors.New("invalid roles")
	}
	tmpInt, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	if tmpInt < 0 || tmpInt > 2 {
		return errors.New("invalid roles")
	}
	*r = Roles(tmpInt)
	return nil
}

type Session interface {
	GetID() int
	GetRole() Roles
}

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"column:name;type:varchar(255);not null;index:idx_user"`
	Email    string `json:"email" gorm:"column:email;type:varchar(255);not null;index:idx_user"`
	Password string `json:"-" gorm:"column:password;type:varchar(255);not null"`
	Roles    Roles  `json:"roles" gorm:"column:roles;type:int;not null;index:idx_user"`
}

func (User) TableName() string {
	return "users"
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetRole() Roles {
	return u.Roles
}

func (u User) CheckPassword(pass string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass))
	if err != nil {
		return false
	}

	return true
}

func (u *User) HashPassword(password string) error {
	var (
		salted        = password
		passwordBytes = []byte(salted)
		cost          = bcrypt.DefaultCost
	)

	hash, errHash := bcrypt.GenerateFromPassword(passwordBytes, cost)
	if errHash != nil {
		return errHash
	}

	u.Password = string(hash)
	return nil
}
