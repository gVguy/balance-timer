package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	if a.session.state == Stopped {
		return false
	}
	response, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         "End session?",
		Message:       "If you close the app, it will end the current session",
		DefaultButton: "Yes",
		CancelButton:  "Cancel",
		Buttons:       []string{"Yes", "Cancel"},
	})
	if err != nil {
		fmt.Printf("beforeClose() Error in MessageDialog, proceeding to close. Error: %v\n", err)
		return false
	}
	fmt.Printf("beforeClose() MessageDialog response: %v\n", response)
	preventClose := response != "Yes"
	return preventClose
}
