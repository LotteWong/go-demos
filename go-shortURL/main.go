package main

import "go-shortURL/router"

func main() {
	app := router.NewApp()
	app.Run()
}
