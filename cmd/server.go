/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	server2 "unicare/internal/server"
)

func startServer() {
	log.Info().Msg("Initializing server...")
	server := server2.New(app)
	port := viper.GetInt("port")
	if port == 0 {
		port = 8080
	}
	go server.Listen(fmt.Sprintf(":%d", port))

	signKill := make(chan os.Signal, 1)
	signal.Notify(signKill, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signKill

	log.Info().Msg("Shutdown server...")
	server.Shutdown()
	log.Info().Msg("Shutdown application...")
	app.Close()
}
