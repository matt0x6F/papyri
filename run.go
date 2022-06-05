package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func RunGUI() error {
	app := NewApp()

	// Create application with options
	return wails.Run(
		&options.App{
			Title:             "Papyri",
			Width:             1024,
			Height:            768,
			DisableResize:     false,
			Fullscreen:        false,
			Frameless:         false,
			MinWidth:          0,
			MinHeight:         0,
			MaxWidth:          0,
			MaxHeight:         0,
			StartHidden:       false,
			HideWindowOnClose: false,
			AlwaysOnTop:       false,
			RGBA:              nil,
			Assets:            assets,
			AssetsHandler:     nil,
			Menu:              app.menu(),
			Logger:            nil,
			LogLevel:          0,
			OnStartup:         app.startup,
			OnDomReady:        nil,
			OnShutdown:        nil,
			OnBeforeClose:     nil,
			Bind: []interface{}{
				app,
			},
			WindowStartState: 0,
			Windows:          nil,
			Mac:              nil,
			Linux:            nil,
		},
	)
}
