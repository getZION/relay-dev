package storage

import (
	v1 "github.com/getzion/relay/gen/proto/zion/v1"
	"github.com/getzion/relay/utils"
	. "github.com/getzion/relay/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
)

var db *gorm.DB

type DBConnectionParams struct {
	Host string `envconfig:"DB_HOST"`
	Name string `envconfig:"DB_NAME"`
	User string `envconfig:"DB_USER"`
	Pass string `envconfig:"DB_PASS"`
}

func ConnectDB() (*gorm.DB, error) {
	var err error

	var params DBConnectionParams
	envconfig.Process("", &params)

	databaseConnectionString := params.User + ":" + params.Pass + "@tcp(" + params.Host + ")/" + params.Name
	Log.Info().Str("databaseConnectionString", databaseConnectionString).Msg("Connecting to database...")

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
