package main

import (
	"github.com/joho/godotenv"
	"github.com/norciTest/notesService/internal/api/service"
	"github.com/norciTest/notesService/internal/config"
	"github.com/norciTest/notesService/internal/storage/pg"
	"log"
	"path/filepath"
)

func main() {
	var (
		storage = &pg.Storage{}
		s       = &service.Service{}
		err     error
	)

	if err = godotenv.Load(filepath.Join(".env")); err != nil {
		log.Fatal(err)
	}

	if err = config.InitConfig(filepath.Join("configs")); err != nil {
		log.Fatal(err)
	}

	if err = storage.Init(); err != nil {
		log.Fatal(err)
	}

	if err = storage.CreateNoteTable(); err != nil {
		log.Fatal(err)
	}

	if s = service.New(storage); err != nil {
		log.Fatal(err)
	}
	if err = s.S.Run(); err != nil {
		log.Fatal(err)
	}
}
