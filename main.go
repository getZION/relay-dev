package main

import (
	"github.com/getzion/relay/api/identityhub"
	"github.com/getzion/relay/api/schema"
	"github.com/getzion/relay/api/storage"
	"github.com/getzion/relay/api/validator"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func main() {

	validator.InitValidator()

	connection, err := storage.NewStorage("mysql")
	if err != nil {
		logrus.Panic(err)
	}

	schemaManager := schema.NewSchemaManager(connection)

	server := identityhub.InitIdentityHubServer(schemaManager)
	logrus.Fatal(fasthttp.ListenAndServe("127.0.0.1:9990", server.Process))
}
