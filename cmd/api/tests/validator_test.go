package main

import (
	"testing"

	"moonlight.zbekxzz.net/internal/validator"
)

func TestValidator_Check(t *testing.T) {
	// Create a new Validator instance
	v := validator.New()

	// Test case 1: When the validation check is successful
	v.Check(true, "key", "message")
	if !v.Valid() {
		t.Error("Expected no errors, but got some.")
	}

	// Test case 2: When the validation check fails
	v.Check(false, "key", "message")
	if v.Valid() {
		t.Error("Expected errors, but got none.")
	}

	// Test case 3: Check if error message is correctly added
	expectedErrorMessage := "message"
	v.Check(false, "key", expectedErrorMessage)
	actualErrorMessage, ok := v.Errors["key"]
	if !ok {
		t.Error("Expected error message not found.")
	}
	if actualErrorMessage != expectedErrorMessage {
		t.Errorf("Expected error message %s, but got %s", expectedErrorMessage, actualErrorMessage)
	}
}
