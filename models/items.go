package models

import "github.com/jinzhu/gorm"

type Item struct {
	Id          int `gorm:primary_key`
	Description string
	Done        bool
}

