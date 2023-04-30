package main

import (
	"fmt"
	"log"

	"github.com/saparaly/forum/internal/handler"
	"github.com/saparaly/forum/internal/repository"
	"github.com/saparaly/forum/internal/server"
	"github.com/saparaly/forum/internal/service"
)

const port = "8000"

func main() {
	db, err := repository.OpenDB("forum.db")
	if err != nil {
		log.Fatalf("error while opening db: %s", err)
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	fmt.Printf("Listening server at post %s", port)
	fmt.Println()
	if err := srv.Run(port, handlers.InitPoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
