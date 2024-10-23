package main

import (
	"context"
)

// App struct
type App struct {
	ctx     context.Context
	session *Session
	config  *Config
	menu    *Menu
}

// NewApp creates a new App application struct
func NewApp() *App {
	app := &App{}
	app.session = NewSession(app)
	app.config = NewConfig().Load()
	app.menu = NewMenu(app)
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.menu.SetMenu()
}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	if a.session.state == Stopped {
		return false
	}
	confirmedStop := a.confirmAction("End session?", "If you close the app, the current session will be stopped")
	return !confirmedStop
}
