/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package main

import (
	"unicare/internal/config"
)

var (
	app *config.App
)

func init() {
	app = config.New(config.Config{})
}
func main() {

}
