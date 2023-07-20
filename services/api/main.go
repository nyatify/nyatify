package main

import (
	"github.com/nyatify/nyatify/pkg/storage"
	"github.com/nyatify/nyatify/services/api/api"
)

func main() {
	db, err := storage.New()
	if err != nil {
		panic(err)
	}

	service := api.NewService(db)

	api.NewServer(service).Run()
}
