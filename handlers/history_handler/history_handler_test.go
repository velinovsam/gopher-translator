package history_handler

import "testing"

func TestUpdateHistory(t *testing.T) {

	original := "bla"
	translation := "blaster"
	UpdateHistory(original, translation)

	if val, ok := history[original]; !ok || val != translation {
		t.Error("Failed to insert entry into history")
	}

}