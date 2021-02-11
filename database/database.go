package database

type AgendaItemModel struct {
	Id          int `gorm:primary_key`
	Description string
	Done        bool
}
