package main

import (
	_ "app/server/conf"
	"app/server/router"
	"log"
)

func main() {
	app := router.New()
	log.Fatal(app.Listen(":30080"))
}
