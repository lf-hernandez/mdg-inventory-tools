package main

import "github.com/lf-hernandez/mdg-inventory-tools/api/config"

func main() {
	app := App{}
	app.Init(config.LoadConfig())
	app.Run()
}
