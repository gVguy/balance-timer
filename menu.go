package main

import (
	"fmt"
	nativeRuntime "runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MenuItemId string

type Menu struct {
	app  *App
	menu *menu.Menu
}

func NewMenu(app *App) *Menu {
	return &Menu{
		app: app,
	}
}

// SetMenu sets app menu to current menu
func (m *Menu) SetMenu() *Menu {
	m.menu = m.createMenu()
	runtime.MenuSetApplicationMenu(m.app.ctx, m.menu)
	runtime.MenuUpdateApplicationMenu(m.app.ctx)
	return m
}

// createMenu creates a new app menu
func (m *Menu) createMenu() *menu.Menu {

	appMenu := menu.NewMenu()

	sessionNew := &menu.MenuItem{
		Type:        "Text",
		Label:       "New",
		Disabled:    m.app.session.stateIs(Stopped),
		Accelerator: keys.CmdOrCtrl("n"),
		Click: func(_ *menu.CallbackData) {
			confirmed := m.app.confirmAction("End session?", "If you proceed to the New Session screen, the current session will be stopped")
			if confirmed {
				m.app.Stop()
			}
		},
	}

	startLabel := "Start"
	if !m.app.session.stateIs(Stopped) {
		startLabel = "Continue"
	}
	sessionStart := &menu.MenuItem{
		Type:        "Text",
		Label:       startLabel,
		Disabled:    m.app.session.stateIs(Working, Resting),
		Accelerator: keys.CmdOrCtrl("g"),
		Click: func(_ *menu.CallbackData) {
			switch m.app.session.state {
			case Stopped:
				fallthrough
			case RestEnd:
				m.app.BeginWork()
			case WorkEnd:
				m.app.BeginRest()
			}
		},
	}

	sessionSkip := &menu.MenuItem{
		Type:        "Text",
		Label:       "Skip period",
		Disabled:    m.app.session.stateIs(Stopped),
		Accelerator: keys.CmdOrCtrl("f"),
		Click: func(_ *menu.CallbackData) {
			thisPeriod := ""
			nextPeriod := ""
			switch m.app.session.state {
			case RestEnd:
				fallthrough
			case Working:
				thisPeriod = "Productivity"
				nextPeriod = "Rest"
			case WorkEnd:
				fallthrough
			case Resting:
				thisPeriod = "Rest"
				nextPeriod = "Productivity"
			}
			confirmTitle := fmt.Sprintf("Skip %v period?", thisPeriod)
			confirmMessage := fmt.Sprintf("The next %v period will start immediately", nextPeriod)
			confirmed := m.app.confirmAction(confirmTitle, confirmMessage)
			if confirmed {
				m.app.Skip()
			}
		},
	}

	sessionMenu := appMenu.AddSubmenu("Session")
	sessionMenu.Append(sessionNew)
	sessionMenu.Append(sessionStart)
	sessionMenu.Append(sessionSkip)

	if nativeRuntime.GOOS == "darwin" {
		// on macOS add "app" menu
		appMenu.Prepend(menu.AppMenu())
		// and also, append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcuts
		appMenu.Append(menu.EditMenu())
	}

	return appMenu
}
