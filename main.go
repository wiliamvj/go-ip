package main

import (
	"go-ip/app"
	"log"
	"os"
)

func main() {

	application := app.Gen()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
