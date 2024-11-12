package main

import (
	"embed"
	"vesgoapp/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	config := backend.Config{
		Host:     "192.168.18.125",
		Port:     3306,
		User:     "**",
		Password: "***",
		Database: "**",
	}

	db := backend.GetConnection(config)
	defer backend.CloseConnection()

	err := wails.Run(&options.App{
		Title:  "vesgoapp",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		Bind: []interface{}{
			db,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
