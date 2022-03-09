package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//go:embed ui/dist
//go:embed ui/dist/_next
//go:embed ui/dist/_next/static/*/*.js
//go:embed ui/dist/_next/static/chunks/*.js
//go:embed ui/dist/_next/static/chunks/pages/*.js
//go:embed ui/dist/_next/static/*/*.js
//go:embed ui/dist/_next/static/css/*.css
var nextFS embed.FS

func serveStaticClient(router *mux.Router) {
	if clientFS, err := fs.Sub(nextFS, "ui/dist"); err != nil {
		logrus.Fatal("failed to load static files subdir")
		return
	} else {
		spaFS := SpaFS{clientFS}
		httpFS := http.FS(spaFS)
		router.PathPrefix("/").Handler(http.FileServer(httpFS))
	}
}

type SpaFS struct {
	base fs.FS
}

func (s SpaFS) Open(name string) (fs.File, error) {
	if file, err := s.base.Open(name); err == nil {
		return file, nil
	} else {
		return s.base.Open("index.html")
	}
}
