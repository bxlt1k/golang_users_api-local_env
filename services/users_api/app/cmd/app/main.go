package main

import (
	"log"
	"users_api/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalln("Unable to start app: " + err.Error())
	}
}
