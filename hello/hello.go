package hello

import "fmt"

func Hello(msg string) string {
	if msg == "" {
		return ""
	}
	return fmt.Sprintf("Hello, %s!!", msg)
}
