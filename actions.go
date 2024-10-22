package main

import (
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// StartSession starts a new session and begins Working period from "Stopped" state with new config
func (a *App) StartSession(workDuration int, restDuration int) {
	fmt.Printf("StartSession() Work duration: %v, Rest duration: %v\n", workDuration, restDuration)
	if !a.session.stateIs(Stopped) {
		return
	}

	a.config.Intervals[Working] = time.Duration(workDuration) * time.Minute
	a.config.Intervals[Resting] = time.Duration(restDuration) * time.Minute
	err := a.config.write()
	if err != nil {
		fmt.Printf("Error in config.write(): %v", err)
	}
	a.BeginWork()
}

// BeginRest begins Rest period from stopped-like state
func (a *App) BeginRest() {
	fmt.Printf("BeginRest()\n")
	if !a.session.stateIs(WorkEnd, RestEnd) {
		return
	}
	a.session.timeRemaining = a.config.Intervals[Resting]
	a.session.setState(Resting)
	go a.session.runSessionPeriod()
}

// BeginWork begins "Working" period with a default duration from stopped-like state
func (a *App) BeginWork() {
	fmt.Println("BeginWork()")
	if !a.session.stateIs(RestEnd, Stopped, WorkEnd) {
		return
	}
	a.beginWork(a.config.Intervals[Working])
}

// MoreWork begins "Working" period with a default duration from WorkEnd state
func (a *App) MoreWork(minutes int) {
	duration := time.Duration(minutes) * time.Minute
	a.session.stopTicker()
	a.beginWork(duration)
}
func (a *App) beginWork(workDuration time.Duration) {
	fmt.Printf("beginWork() Work duration: %v\n", workDuration)
	a.session.timeRemaining = workDuration
	a.session.setState(Working)
	go a.session.runSessionPeriod()
	runtime.WindowUnfullscreen(a.ctx)
}

// Stop resets to the "Stopped" state from any other state
func (a *App) Stop() {
	fmt.Println("Stop()")
	if a.session.stateIs(Stopped) {
		return
	}
	a.session.stopTicker()
	a.session.setState(Stopped)
	runtime.WindowUnfullscreen(a.ctx)
}

// SkipWork skips the current "Working" period to "Resting" mode
func (a *App) SkipWork() {
	fmt.Println("SkipWork()")
	if !a.session.stateIs(Working) {
		return
	}
	a.session.stopTicker()
	a.session.setState(WorkEnd)
	a.BeginRest()
}

// SkipRest skips the current "Resting" period to "Working" mode
func (a *App) SkipRest() {
	fmt.Println("SkipRest()")
	if !a.session.stateIs(Resting) {
		return
	}
	a.session.stopTicker()
	a.session.setState(RestEnd)
	a.BeginWork()
}

// Returns config with intervals formatted to minutes
func (a *App) GetConfigMinutes() Config {
	return Config{
		Intervals: IntervalsMap{
			Working: a.config.Intervals[Working] / time.Minute,
			Resting: a.config.Intervals[Resting] / time.Minute,
		},
	}
}

// helper function to check whether the current state of the session matches the required
func (s *Session) stateIs(allowedStates ...SessionState) bool {
	for _, state := range allowedStates {
		if s.state == state {
			return true
		}
	}
	fmt.Printf("stateIs() Expected state(s): %v. Current: '%v')\n", allowedStates, s.state)
	return false
}

// // helper function to round float minutes to time duration of int number of seconds
// func roundToSec(min float32) time.Duration {
// 	minutes := min * float32(time.Minute)
// 	roundedSeconds := int(minutes / float32(time.Second))
// 	return time.Duration(roundedSeconds * int(time.Second))
// }
