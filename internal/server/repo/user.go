/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package repo

import "unicare/internal/model"

type UserInterface interface {
	GetUser(id int) (*model.User, error)
}
