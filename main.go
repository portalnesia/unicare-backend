/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package main

import (
	"unicare/cmd"
	"unicare/internal/config"
)

var build string

func main() {
	cfg := config.Config{
		Build: build,
	}
	cmd.Start(cfg)
}
