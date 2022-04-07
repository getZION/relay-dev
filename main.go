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

	connection, err := storage.NewStorage("mysql")
	if err != nil {
		logrus.Panic(err)
	}

	schemaManager := schema.NewSchemaManager(connection)
	server := identityhub.InitIdentityHubServer(schemaManager)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	logrus.Fatal(server.Listen(fmt.Sprintf("%s:%s", host, port)))
}
