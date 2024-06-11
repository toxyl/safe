package safe

import (
	"testing"
)

func TestSafeFloat_Add(t *testing.T) {
	tests := []struct {
		name     string
		sf       *SafeFloat
		d        float64
		expected float64
	}{
		{"test 1", NewSafeFloat(10.0), 5.1, 15.1},
		{"test 2", NewSafeFloat(10.0), 32.43, 42.43},
		{"test 3", NewSafeFloat(10.0), -5.7, 4.3},
		{"test 4", NewSafeFloat(10.0), -11.5, -1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sf.Add(tt.d)
			if val := tt.sf.Get(); val != tt.expected {
				t.Errorf("Expected %f, got %f", tt.expected, val)
			} else {
				t.Logf("%f matches expectation", val)
			}

		})
	}
}
