package main

import (
	"context"
)

// App struct
type App struct {
	ctx     context.Context
	session *Session
	config  *Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	app := &App{}
	app.session = NewSession(app)
	app.config = NewConfig().load()
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
