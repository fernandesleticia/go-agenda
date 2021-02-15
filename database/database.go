package database

type Item struct {
	Id          int `gorm:primary_key`
	Description string
	Done        bool
}
