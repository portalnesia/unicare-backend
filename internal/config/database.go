/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package config

import (
	"fmt"
	mysql2 "github.com/go-sql-driver/mysql"
	"github.com/golang-module/carbon"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func initDatabase() (db *gorm.DB) {
	var err error
	location, errLoc := time.LoadLocation("Asia/Jakarta")
	if errLoc != nil {
		panic(errLoc)
	}

	mysqlconfig := mysql2.Config{
		User:                 viper.GetString("db.mysql.user"),
		Passwd:               viper.GetString("db.mysql.password"),
		DBName:               viper.GetString("db.mysql.database"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", viper.GetString("db.mysql.host"), viper.GetInt("db.mysql.port")),
		ParseTime:            true,
		Loc:                  location,
		AllowNativePasswords: true,
	}
	dsn := mysqlconfig.FormatDSN()

	portalnesia := mysql.New(mysql.Config{
		DSN: dsn,
	})

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err = gorm.Open(portalnesia, &gorm.Config{
		PrepareStmt: true,
		Logger:      newLogger,
		NowFunc: func() time.Time {
			return carbon.Now("Asia/Jakarta").ToStdTime()
		},
	})

	if err != nil {
		panic(err)
	}

	return
}
