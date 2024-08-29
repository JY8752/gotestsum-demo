package hello_test

import (
	"gotestsum-demo/hello"
	"testing"
)

func TestHello(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		msg      string
		expected string
	}{
		"Hello, gotestsum!!": {
			msg:      "gotestsum",
			expected: "Hello, gotestsum!!",
		},
		"Hello, World!!": {
			msg:      "World",
			expected: "Hello, World!!",
		},
		"Hello, Go!!": {
			msg:      "Go",
			expected: "Hello, Go!!",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if result := hello.Hello(tt.msg); result != tt.expected {
				t.Errorf("expected %s, but %s\n", tt.expected, result)
			}
		})
	}
}

func TestHello2(t *testing.T) {
	if result := hello.Hello(""); result != "" {
		t.Errorf("expected empty, but %s\n", result)
	}
}
