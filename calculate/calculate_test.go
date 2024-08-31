package calculate_test

import (
	"gotestsum-demo/calculate"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		x, y, expected int
	}{
		"1+1":   {x: 1, y: 1, expected: 2},
		"2+1":   {x: 2, y: 1, expected: 3},
		"1-1":   {x: 1, y: -1, expected: 0},
		"1+2":   {x: 1, y: 2, expected: 3},
		"1+0":   {x: 1, y: 0, expected: 1},
		"10+10": {x: 10, y: 10, expected: 20},
		// "10+1":  {x: 10, y: 1, expected: 21},
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

func Test10S(t *testing.T) {
	t.Parallel()
	time.Sleep(10 * time.Second)
}

func Test5S(t *testing.T) {
	t.Parallel()
	time.Sleep(5 * time.Second)
}

func Test3S(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
