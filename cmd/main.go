/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"unicare/internal/config"
	server2 "unicare/internal/server"
)

// Start
//
//	@Description: Start server
//	@param cfg
func Start(cfg config.Config) {
	app := config.New(cfg)
	server := server2.New(app)
	port := viper.GetInt("port")
	if port == 0 {
		port = 8080
	}
	go server.Listen(fmt.Sprintf(":%d", port))

	signKill := make(chan os.Signal, 1)
	signal.Notify(signKill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signKill

	server.Shutdown()
	app.Close()
}
