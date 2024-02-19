/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

type Env struct {
	IsProd bool
}

var env Env

type Config struct {
	Build string
}

type App struct {
	Env       Env
	DB        *gorm.DB
	Exc       Exception
	Validator *validator.Validate
}

func New(cfg Config) *App {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	log.Info().Str("Build", cfg.Build).Msg("Initializing application...")
	_ = gotenv.Load()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	//_ = viper.ReadInConfig()

	env = Env{
		IsProd: cfg.Build == "production",
	}

	return &App{
		Env:       env,
		Exc:       Exception{},
		DB:        initDatabase(),
		Validator: initValidate(),
	}
}

func (a *App) Close() {
	if a.DB != nil {
		if db, err := a.DB.DB(); err == nil {
			db.Close()
		}
	}
}
