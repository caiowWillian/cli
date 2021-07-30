package main

import (
	"fmt"
	"log"
	"module/src/app"
	"os"
)

func main() {
	fmt.Println("start")
	application := app.Start()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
