package safe

import (
	"testing"
	"time"
)

func TestSafeDuration_Add(t *testing.T) {
	tests := []struct {
		name     string
		sf       *SafeDuration
		d        time.Duration
		expected time.Duration
	}{
		{"test 1", NewSafeDuration(10 * time.Second), 5 * time.Second, 15 * time.Second},
		{"test 2", NewSafeDuration(10 * time.Hour), 5 * time.Minute, 10*time.Hour + 5*time.Minute},
		{"test 3", NewSafeDuration(10 * time.Hour), -5 * time.Minute, 9*time.Hour + 55*time.Minute},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sf.Add(tt.d)
			if val := tt.sf.Get(); val != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected.String(), val.String())
			} else {
				t.Logf("%s matches expectation", val.String())
			}

		})
	}
}
