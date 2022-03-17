package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type mysqlConnectionParams struct {
	Host string `envconfig:"DB_HOST"`
	Name string `envconfig:"DB_NAME"`
	User string `envconfig:"DB_USER"`
	Pass string `envconfig:"DB_PASS"`
}

func NewMySqlConnection() (*gorm.DB, error) {

	var err error
	var db *gorm.DB
	var params mysqlConnectionParams
	envconfig.Process("", &params)

	databaseConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", params.User, params.Pass, params.Host, params.Name)
	logrus.Infof("Connecting to MySQL database: %s", databaseConnectionString)

	db, err = gorm.Open("mysql", databaseConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
