package main

import (
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) emitState(state SessionState) {
	fmt.Printf("emitState() Emit session state update: '%v'\n", state)
	runtime.EventsEmit(a.ctx, "state", state)
}

func (a *App) emitTickerStart(t time.Duration) {
	fmt.Printf("emitState() Emit ticker start: '%v'\n", t)
	runtime.EventsEmit(a.ctx, "tickerStart", t.Seconds(), t.String())
}

func (a *App) emitTime(t time.Duration) {
	// fmt.Printf("emitTime() Emit session remaining time update: %v", a.session.timeRemaining)
	runtime.EventsEmit(a.ctx, "time", t.Seconds(), t.String())
}
