package lib

import (
	"fmt"
	"os"
)

var PrintFunc func(msg string, args ...any) (int, error) = fmt.Printf

var ExitFunc func() = func() {
	os.Exit(0)
}

func ThrowError(msg string, reason error) {
	PrintFunc("Error: %s\nReason: %v\n", msg, reason)
	ExitFunc()
}
