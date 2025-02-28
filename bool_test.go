package safe

import (
	"testing"
)

func TestSafeBool(t *testing.T) {
	tests := []struct {
		name            string
		sf              *SafeBool
		expectedToggle  bool
		expectedEnable  bool
		expectedDisable bool
	}{
		{"test 1", NewSafeBool(true), false, true, false},
		{"test 2", NewSafeBool(false), true, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if bToggle := tt.sf.Toggle(); bToggle != tt.expectedToggle {
				t.Errorf("toggle: expected %v, got %v", tt.expectedToggle, bToggle)
			} else {
				t.Logf("toggle: %v matches expectation", bToggle)
			}

			if bEnable := tt.sf.Enable(); bEnable != tt.expectedEnable {
				t.Errorf("enable: expected %v, got %v", tt.expectedEnable, bEnable)
			} else {
				t.Logf("enable: %v matches expectation", bEnable)
			}

			if bDisable := tt.sf.Disable(); bDisable != tt.expectedDisable {
				t.Errorf("disable: expected %v, got %v", tt.expectedDisable, bDisable)
			} else {
				t.Logf("disable: %v matches expectation", bDisable)
			}
		})
	}
}
