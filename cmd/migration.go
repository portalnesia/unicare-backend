/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package cmd

import (
	"unicare/internal/model"
)

func startMgration() {
	app.DB.AutoMigrate(
		model.User{},
		model.Customer{},
	)
}
