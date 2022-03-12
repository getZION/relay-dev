package collections

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getzion/relay/api"
	"github.com/getzion/relay/api/datastore"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

const (
	OBJECT_ID           = "d82c0026-ed42-4b26-81f3-94805958a75c"
	INVALID             = "<invalid>"
	DATE_CREATED        = "1645917431"
	DATE_PUBLISHED      = "1645917431"
	SCHEMA_ORGANIZATION = "https://schema.org/Organization"
	SCHEMA_CONVERSATION = "https://schema.org/Conversation"
	SCHEMA_MESSAGE      = "https://schema.org/SocialMediaPosting"
	SCHEMA_PAYMENT      = "https://schema.org/Invoice"
	SCHEMA_PERSON       = "https://schema.org/Person"
)

func initTestDataStore() (*datastore.Store, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		logrus.Panic(err)
	}

	gormDb, err := gorm.Open("mysql", db)
	if err != nil {
		logrus.Panic(err)
	}

	connection := &api.Connection{
		DB: gormDb,
	}

	store, err := datastore.NewStore(connection)
	if err != nil {
		logrus.Panic(err)
	}

	return store, mock
}
