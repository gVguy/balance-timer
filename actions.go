package main

import (
	"fmt"
	"time"
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

// BeginRest begins Rest period from "WorkEnd" state
func (a *App) BeginRest() {
	fmt.Printf("BeginRest()\n")
	if !a.session.stateIs(WorkEnd) {
		return
	}
	a.session.timeRemaining = a.config.Intervals[Resting]
	a.session.setState(Resting)
	go a.session.runSessionPeriod()
}

// BeginWork begins Work period from "RestEnd" or "Stopped" state
func (a *App) BeginWork() {
	fmt.Printf("BeginWork()\n")
	if !a.session.stateIs(RestEnd, Stopped) {
		return
	}
	a.session.timeRemaining = a.config.Intervals[Working]
	a.session.setState(Working)
	go a.session.runSessionPeriod()
}

// Stop resets to the "Stopped" state from any other state
func (a *App) Stop() {
	fmt.Println("Stop()")
	if a.session.stateIs(Stopped) {
		return
	}
	select {
	case a.session.stopChan <- true:
		fmt.Println("sent a stop signal to the Stop chanel")
	default:
		fmt.Println(">> skipped sending a stop signal to the Stop chanel (no receiver)")
	}
	a.session.timeRemaining = 0
	a.session.setState(Stopped)
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
