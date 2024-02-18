/*
 * Copyright (c) Northbit - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written by Putu Aditya <aditya@portalnesia.com>
 */

package entity

import (
	"go.portalnesia.com/nullable"
	"math"
	"strings"
)

type Pagination struct {
	Page     int    `query:"page"`
	Total    int64  `query:"-"`
	PageSize int    `query:"page_size"`
	Start    int    `query:"-"`
	Q        string `query:"q"`

	Active nullable.Bool `query:"active"`
	done   bool          `query:"-"`
}

func (d *Pagination) Init() {
	if d.done {
		return
	}

	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 || d.PageSize > 50 {
		d.PageSize = 15
	}

	if d.Q != "" {
		d.Q = strings.ToLower(d.Q)
	}

	start := 0
	if d.Page > 1 {
		start = int(d.Page)*d.PageSize - d.PageSize
	}
	d.Start = start

	d.done = true
}

type ResponseDataPaginationData[T any] struct {
	Page      int   `json:"page"`
	TotalPage int   `json:"total_page"`
	Total     int64 `json:"total"`
	Data      T     `json:"data"`
}

func (d *Pagination) NewPagination(data any) ResponseDataPaginationData[any] {
	totalPage := int(math.Ceil(float64(d.Total) / float64(d.PageSize)))

	return ResponseDataPaginationData[any]{
		Page:      d.Page,
		Total:     d.Total,
		TotalPage: totalPage,
		Data:      data,
	}
}
