package database

import "github.com/jinzhu/gorm"

var MysqlInstance, _ = gorm.Open("mysql", "root:root@/agenda?charset=utf8&parseTime=True&loc=Local")

type ItemHandler interface {
	CreateItem(description string, done bool) ([]*models.Item, error)
	UpdateItem(id int, done bool) ([]*models.Item, error)
	DeleteItem(id int) ([]*models.Item, error)
	GetDoneItems() (bool, error)
	GetPendingItems() (bool, error)
}