package maxchars

import (
	"testing"
)

func TestMaxChars(t *testing.T) {
	for _, tt := range testCases {
		res := maxChars(tt.input)

		if res != tt.expected {
			t.Errorf("Expected %s but got %s", tt.expected, res)
			return
		}

		t.Logf("Expected %s and got %s", tt.expected, res)
	}
}
