/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package model

import "go.portalnesia.com/nullable"

type Context struct {
	User     nullable.Type[User]
	Customer nullable.Type[Customer]

	IP     string
	Method string
}
