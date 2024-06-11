package safe

import (
	"testing"
)

func TestSafeUint_Add(t *testing.T) {
	tests := []struct {
		name     string
		sf       *SafeUint
		d        int
		expected uint
	}{
		{"test 1", NewSafeUint(uint(10)), 5, 15},
		{"test 2", NewSafeUint(uint(10)), 32, 42},
		{"test 3", NewSafeUint(uint(10)), -5, 5},
		{"test 4", NewSafeUint(uint(10)), -11, 0},
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
