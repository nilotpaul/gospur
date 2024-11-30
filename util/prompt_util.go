package util

import (
	"fmt"

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
	items := ms.Items
	selected := make(map[int]bool)
	cursor := 0

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

func getSelectedItems(items []string, selected map[int]bool) []string {
	selectedItems := make([]string, 0)
	for i, isSelected := range selected {
		if isSelected {
			selectedItems = append(selectedItems, items[i])
		}
	}

	return selectedItems
}
