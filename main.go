package main

import (
	"fmt"
	"os"

	"github.com/getzion/relay/api/identityhub"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	"github.com/sirupsen/logrus"
)

func main() {

	validator.InitValidator()

	storage, err := storage.NewStorage("mysql")
	if err != nil {
		logrus.Panic(err)
	}

	schemaManager := schema.NewSchemaManager(storage)
	server := identityhub.InitIdentityHubServer(schemaManager, storage)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	logrus.Fatal(server.Listen(fmt.Sprintf("%s:%s", host, port)))
}
