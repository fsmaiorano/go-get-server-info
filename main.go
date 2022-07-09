package main

import (
	"go-get-server-info/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
