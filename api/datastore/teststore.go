package datastore

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/validator"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func NewTestStore() (*Store, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		logrus.Panic(err)
	}

	gormDb, err := gorm.Open("mysql", db)
	if err != nil {
		logrus.Panic(err)
	}

	gormDb.LogMode(true)

	connection := &api.Connection{
		DB: gormDb,
	}

	store, err := NewStore(connection)
	if err != nil {
		logrus.Panic(err)
	}

	//todo: move something different
	validator.InitValidator()
	return store, mock
}
