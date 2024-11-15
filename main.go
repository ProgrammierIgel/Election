package main

import (
	"fmt"
	"net/http"

	"github.com/programmierigel/voting/api"
	"github.com/programmierigel/voting/environment"
	"github.com/programmierigel/voting/storage/inmemory"
)

func main() {
	port, err := environment.Port(3000)
	if err != nil {
		panic(err)
	}

	password, err := environment.Password()
	if err != nil {
		panic(err)
	}

	store := inmemory.New(password)

	router := api.GetRouter(store, password)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
