package main

import (
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type SessionState string

const (
	Stopped SessionState = "Stopped"
	Working SessionState = "Working"
	WorkEnd SessionState = "WorkEnd"
	Resting SessionState = "Resting"
	RestEnd SessionState = "RestEnd"
)

type Session struct {
	state         SessionState
	stopChan      chan bool
	timeRemaining time.Duration
	app           *App
}

func NewSession(app *App) *Session {
	return &Session{
		state:    Stopped,
		app:      app,
		stopChan: make(chan bool),
	}
}

func (s *Session) setState(state SessionState) {
	fmt.Printf("setState() '%s'\n", state)
	s.state = state
	s.app.emitState(state)
}

func (s *Session) runSessionPeriod() {
	fmt.Println("runSessionPeriod() Starting the ticker...")
	ticker := time.NewTicker(1 * time.Second)
	s.app.emitTickerStart(s.timeRemaining)
	defer ticker.Stop()
	defer fmt.Println("runSessionPeriod() done. Deferred ticker stopped")
	for {
		select {
		case <-s.stopChan:
			fmt.Println("runSessionPeriod() Received an event in Stop chanel")
			return
		case <-ticker.C:
			s.timeRemaining -= 1 * time.Second
			s.app.emitTime(s.timeRemaining)
			if s.timeRemaining <= 0 {
				fmt.Println("runSessionPeriod() Time's up")
				s.handleTimeEnd()
				return
			}
		}
	}
}

func (s *Session) handleTimeEnd() {
	fmt.Printf("handleTimeEnd() Current state: %v\n", s.state)
	switch s.state {
	case Working: // finished working
		runtime.WindowFullscreen(s.app.ctx)
		s.setState(WorkEnd)
		// run new ticker for auto-start rest
		s.timeRemaining = s.app.config.Intervals[WorkEnd]
		s.runSessionPeriod()
	case Resting: // finished resting
		runtime.WindowFullscreen(s.app.ctx)
		s.setState(RestEnd)
	case WorkEnd: // auto-continue ticker finished
		s.app.BeginRest()
	}
}

func (s *Session) stopTicker() {
	fmt.Println("stopTicker()")
	select {
	case s.stopChan <- true:
		fmt.Println("sent a stop signal to the Stop chanel")
	default:
		fmt.Println(">> skipped sending a stop signal to the Stop chanel (no receiver)")
	}
	s.timeRemaining = 0
}
