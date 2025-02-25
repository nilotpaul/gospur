package ui

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

var done = doneColor(promptui.IconGood)

type Spinner struct {
	// Message to show beside the loading icon
	loadingMsg string

	frames []string
	// Spin delay
	delay time.Duration
	// Channel for stopping the spinner
	stopChan chan struct{}
}

func NewSpinner(msg string) *Spinner {
	return &Spinner{
		loadingMsg: msg,
		frames:     []string{"|", "/", "-", "\\"},
		delay:      100 * time.Millisecond,
		stopChan:   make(chan struct{}),
	}
}

func (s *Spinner) Start() {
	go func() {
		i := 0
		for {
			select {
			case <-s.stopChan:
				fmt.Printf("\r%s\n", done)
				return

			default:
				spin := promptui.Styler(promptui.FGBlue)(s.frames[i%len(s.frames)])
				fmt.Printf("\r%s %s", spin, s.loadingMsg)
				time.Sleep(s.delay)
				i++
			}

		}
	}()
}

func (s *Spinner) Stop() {
	close(s.stopChan)
	// Waiting for 100ms to keep the stdout synchronised.
	time.Sleep(100 * time.Millisecond)
}
