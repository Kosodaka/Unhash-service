package main

import (
	"log"
	"unhashService/app"
	"unhashService/entity"
)

func main() {

	cfg, err := entity.NewConfig("./.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("%+v\n", cfg)
	app, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}
	app.Run()
}
