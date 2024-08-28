package calculate_test

import (
	"gotestsum-demo/calculate"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		x, y, expected int
	}{
		"1+1": {x: 1, y: 1, expected: 2},
		"2+1": {x: 2, y: 1, expected: 3},
		"1-1": {x: 1, y: -1, expected: 0},
		"1+2": {x: 1, y: 2, expected: 3},
		// "10+1": {x: 10, y: 1, expected: 21},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if result := calculate.Add(tt.x, tt.y); result != tt.expected {
				t.Errorf("expected %d, but %d\n", tt.expected, result)
			}
		})
	}
}
