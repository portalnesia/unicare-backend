/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package server

import (
	"github.com/gofiber/fiber/v2"
	"go.portalnesia.com/nullable"
	"reflect"
	"unicare/internal/model"
)

func registerDecoder() []fiber.ParserType {

	nullBool := fiber.ParserType{
		Customtype: nullable.Bool{},
		Converter:  nullable.Bool{}.FiberConverter,
	}
	nullTime := fiber.ParserType{
		Customtype: nullable.Time{},
		Converter:  nullable.Time{}.FiberConverter,
	}
	nullString := fiber.ParserType{
		Customtype: nullable.String{},
		Converter:  nullable.String{}.FiberConverter,
	}
	nullFloat := fiber.ParserType{
		Customtype: nullable.Float{},
		Converter:  nullable.Float{}.FiberConverter,
	}
	nullInt := fiber.ParserType{
		Customtype: nullable.Int{},
		Converter:  nullable.Int{}.FiberConverter,
	}

	nullStringArray := fiber.ParserType{
		Customtype: nullable.StringArray{},
		Converter:  nullable.StringArray{}.FiberConverter,
	}

	rolesType := fiber.ParserType{
		Customtype: model.Roles(0),
		Converter:  RoleDecoder,
	}

	return []fiber.ParserType{
		nullBool,
		nullTime,
		nullString,
		nullFloat,
		nullInt,
		nullInt,
		nullStringArray,
		rolesType,
	}
}

func RoleDecoder(s string) reflect.Value {
	var d model.Roles
	if r, ok := model.StringToRoles[s]; ok {
		d = r
	}
	return reflect.ValueOf(d)
}
