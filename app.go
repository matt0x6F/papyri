package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ToggleAboutModal() {
	runtime.EventsEmit(a.ctx, "toggle-about", true)
}

func (a *App) Test() {

}

func (a *App) menu() *menu.Menu {
	parent := menu.NewMenu()

	// app menu
	app := parent.AddSubmenu("Papyri")
	app.AddText(
		"Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
			a.Quit()
		},
	)

	// help menu
	help := parent.AddSubmenu("Help")
	help.AddText(
		"About", keys.CmdOrCtrl("?"), func(data *menu.CallbackData) {
			a.ToggleAboutModal()
		},
	)

	return parent
}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
}
