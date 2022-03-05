package storage

import (
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/getzion/relay/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectSQL(databaseConnectionString string) (*gorm.DB, error) {
	var err error

	db, err = gorm.Open("mysql", databaseConnectionString)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&v1.CommunityORM{},
		&v1.ContactORM{},
		&v1.ConversationORM{},
		&v1.CommentORM{},
		&v1.MessageORM{},
		&v1.InvoiceORM{},
		&v1.UserORM{},
	)

	utils.Log.Info().Msg("Connected to MySQL database. Migrations successful.")
	return db, nil
}
