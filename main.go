package main

import (
	"fmt"
	"go-ip/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Init")

	application := app.Gen()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
