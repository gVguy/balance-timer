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
	err := a.config.Write()
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

// Skip handles the skip logic based on current period and calls the corresponding action
func (a *App) Skip() {
	fmt.Println("Skip()")
	if a.session.stateIs(Stopped) {
		return
	}
	switch a.session.state {
	case Working:
		a.SkipWork()
	case Resting:
		a.SkipRest()
	case WorkEnd: // in work end, skip rest
		a.BeginWork()
	case RestEnd: // in rest end, skip work
		a.BeginRest()
	}
}

// GetConfigMinutes returns config with intervals formatted to minutes
func (a *App) GetConfigMinutes() Config {
	return Config{
		Intervals: IntervalsMap{
			Working: a.config.Intervals[Working] / time.Minute,
			Resting: a.config.Intervals[Resting] / time.Minute,
		},
	}
}

// stateIs checks whether the current state of the session matches the required
func (s *Session) stateIs(allowedStates ...SessionState) bool {
	for _, state := range allowedStates {
		if s.state == state {
			return true
		}
	}
	// fmt.Printf("stateIs() Expected state(s): %v. Current: '%v')\n", allowedStates, s.state)
	return false
}

// confirmAction prompts user for confirmation with a dialog
func (a *App) confirmAction(title string, message string) bool {
	response, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         title,
		Message:       message,
		DefaultButton: "Yes",
		CancelButton:  "Cancel",
		Buttons:       []string{"Yes", "Cancel"},
	})
	if err != nil {
		fmt.Printf("confirmAction() Error in MessageDialog, treat as confirmed. Error: %v\n", err)
		return true
	}
	fmt.Printf("confirmAction() MessageDialog response: %v\n", response)
	confirmedStop := response == "Yes"
	return confirmedStop
}
