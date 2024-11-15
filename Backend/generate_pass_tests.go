package main

import (
	"testing"
)

// Test generatePassword function
func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		length int
	}{
		{length: 0},  // Default length (12)
		{length: 5},  // Short password
		{length: 20}, // Long password
	}

	for _, tt := range tests {
		t.Run(string(tt.length), func(t *testing.T) {
			password := generatePassword(tt.length)

			// Check the length of the generated password
			expectedLength := tt.length
			if tt.length <= 0 {
				expectedLength = 12 // Default length
			}
			if len(password) != expectedLength {
				t.Errorf("Expected length %d, got %d", expectedLength, len(password))
			}

			// Check if the password is not empty
			if password == "" {
				t.Error("Generated password is empty")
			}
		})
	}
}

// Test handlera Lambda
func TestHandler(t *testing.T) {
	tests := []struct {
		request  Request
		expected int
	}{
		{request: Request{Length: 0}, expected: 12},  // Default length (12)
		{request: Request{Length: 8}, expected: 8},   // Short password
		{request: Request{Length: 16}, expected: 16}, // Long password
	}

	for _, tt := range tests {
		t.Run(string(tt.request.Length), func(t *testing.T) {
			response, err := handler(nil, tt.request)
			if err != nil {
				t.Errorf("Handler returned an error: %v", err)
			}

			// Check the length of the generated password
			if len(response.Password) != tt.expected {
				t.Errorf("Expected length %d, got %d", tt.expected, len(response.Password))
			}
		})
	}
}
