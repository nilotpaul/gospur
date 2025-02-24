package util

import (
	"fmt"
	"time"

	"github.com/manifoldco/promptui"
)

var doneColor = promptui.Styler(promptui.FGGreen)

type MultiSelect struct {
	Label string
	Items []string

	// Size is the number of items that should appear on the select before scrolling is necessary.
	// Defaults to 5.
	Size int
}

// MultiSelect provides a prompt for selecting multiple items from a list of strings.
func (ms MultiSelect) Run() ([]string, error) {
	// Track selections as a map for O(1) updates
	var (
		items    = ms.Items
		selected = make(map[int]bool)
		cursor   = 0
	)

	for {
		// Prepare items with bullets and checkmarks for rendering
		displayItems := make([]string, len(items)+1) // +1 for the "Done" button
		for i, item := range items {
			if selected[i] {
				displayItems[i] = fmt.Sprintf("%s %s", promptui.IconGood, item) // Selected
			} else {
				displayItems[i] = item // Unselected
			}
		}
		displayItems[len(items)] = doneColor("Click here to continue")

		prompt := promptui.Select{
			Label:        ms.Label,
			Items:        displayItems,
			Size:         ms.Size,
			CursorPos:    cursor,
			HideSelected: true,
		}

		// Run the prompt
		index, _, err := prompt.Run()
		if err != nil {
			return nil, fmt.Errorf("prompt failed: %v", err)
		}

		// If the "Done" button is selected, exit the loop.
		if index == len(items) {
			break
		}

		// Toggle selection state
		selected[index] = !selected[index]

		// Remember the last cursor position
		cursor = index
	}

	// Collect selected items
	selectedItems := getSelectedItems(items, selected)

	return selectedItems, nil
}

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
				fmt.Printf("\r%s\n", promptui.Styler(promptui.FGGreen)("✔"))
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

func getSelectedItems(items []string, selected map[int]bool) []string {
	selectedItems := make([]string, 0)
	for i, isSelected := range selected {
		if isSelected {
			selectedItems = append(selectedItems, items[i])
		}
	}

	return selectedItems
}
