/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package model

import "time"

type BaseModel struct {
	ID        int        `json:"id" gorm:"column:id;primaryKey;autoIncrement;type:int"`
	CreatedAt time.Time  `json:"-" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"-" gorm:"column:updated_at;autoUpdateTime"`
}
