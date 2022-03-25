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
	db.LogMode(true)

	db.AutoMigrate(
		&v1.CommunityORM{},
		//&v1.ContactORM{},
		&v1.ConversationORM{},
		&v1.CommentORM{},
		&v1.MessageORM{},
		//&v1.InvoiceORM{},
		&v1.UserORM{},
		&v1.Tag{},
	)

	db.Table("community_users").RemoveIndex("CommunityZid").AddIndex("CommunityZid", "CommunityZid")
	db.Table("community_users").RemoveIndex("UserDid").AddIndex("UserDid", "UserDid")
	db.Table("community_users").AddForeignKey("CommunityZid", "communities(Zid)", "RESTRICT", "RESTRICT")
	db.Table("community_users").AddForeignKey("UserDid", "users(Did)", "RESTRICT", "RESTRICT")

	db.Table("community_tags").RemoveIndex("CommunityId").AddIndex("CommunityId", "CommunityId")
	db.Table("community_tags").RemoveIndex("TagId").AddIndex("TagId", "TagId")
	db.Table("community_tags").AddForeignKey("CommunityId", "communities(Id)", "RESTRICT", "RESTRICT")
	db.Table("community_tags").AddForeignKey("TagId", "tags(Id)", "RESTRICT", "RESTRICT")

	db.Table("conversations").AddIndex("community_zid", "community_zid")
	db.Table("conversations").AddForeignKey("community_zid", "communities(Zid)", "RESTRICT", "RESTRICT")

	db.Table("comments").AddIndex("conversation_zid", "conversation_zid")
	db.Table("comments").AddIndex("user_did", "user_did")
	db.Table("comments").AddForeignKey("conversation_zid", "conversations(Zid)", "RESTRICT", "RESTRICT")
	db.Table("comments").AddForeignKey("user_did", "users(Did)", "RESTRICT", "RESTRICT")

	logrus.Info("Migrations successful.")
	return api.InitDatabase(db), nil
}
