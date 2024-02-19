/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */
package cmd

import (
	"flag"
	"unicare/internal/config"
)

var (
	app *config.App
)

// Start
//
//	@Description: Start server
//	@param cfg
func Start(cfg config.Config) {
	app = config.New(cfg)

	migration := flag.Bool("migration", false, "Database Migration")
	flag.Parse()

	if *migration {
		startMgration()
	} else {
		startServer()
	}
}
