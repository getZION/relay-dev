package api

import "github.com/jinzhu/gorm"

//todo: maybe in the future we can consider about leaving use gorm and use native libraries. so I leave it not encapsulated
type Connection struct {
	DB *gorm.DB
}

func InitDatabase(database *gorm.DB) *Connection {
	return &Connection{
		DB: database,
	}
}
