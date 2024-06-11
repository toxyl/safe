package safe

import (
	"testing"
)

func TestSafeInt_Add(t *testing.T) {
	tests := []struct {
		name     string
		sf       *SafeInt
		d        int
		expected int
	}{
		{"test 1", NewSafeInt(10), 5, 15},
		{"test 2", NewSafeInt(10), 32, 42},
		{"test 3", NewSafeInt(10), -5, 5},
		{"test 4", NewSafeInt(10), -11, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sf.Add(tt.d)
			if val := tt.sf.Get(); val != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, val)
			} else {
				t.Logf("%d matches expectation", val)
			}

		})
	}
}
