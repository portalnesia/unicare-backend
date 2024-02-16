/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package config

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"gorm.io/gorm"
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
	Env Env
	DB  *gorm.DB
	Exc Exception
}

func New(cfg Config) *App {
	_ = gotenv.Load()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	env = Env{
		IsProd: cfg.Build == "production",
	}
	zerolog.TimeFieldFormat = time.RFC3339

	return &App{
		Env: env,
		Exc: Exception{},
		//DB:  initDatabase(),
	}
}

func (a *App) Close() {
	if a.DB != nil {
		if db, err := a.DB.DB(); err == nil {
			db.Close()
		}
	}
}
