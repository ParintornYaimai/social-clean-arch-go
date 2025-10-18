package main

import (
	"log"

	"github.com/ParintornYaimai/socialmedia-go/internal/env"
	"github.com/ParintornYaimai/socialmedia-go/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewPostgrestStorage(nil)
	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
