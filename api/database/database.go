package database

import (
	"fmt"

	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/database/mysql"
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// NewDatabase should use config options to return a connection to the requested database
func NewDatabase(storeType string) (connection *api.Connection, err error) {

	var db *gorm.DB
	switch storeType {
	case "mysql":
		db, err = mysql.NewMySqlConnection()
	default:
		return nil, fmt.Errorf("unknown storage database: %s", storeType)
	}

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

	logrus.Info("Migrations successful.")
	return api.InitDatabase(db), nil
}
