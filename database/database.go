package database

import "github.com/jinzhu/gorm"

var MysqlInstance, _ = gorm.Open("mysql", "root:root@/agenda?charset=utf8&parseTime=True&loc=Local")
