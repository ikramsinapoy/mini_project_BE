package main

import (
	"foodcal/configs"
	"foodcal/routes"
)

func main() {
	configs.ConnectDB()
	e := routes.New()
	e.Start(":8000")
}
