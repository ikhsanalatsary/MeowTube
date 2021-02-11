package logger

import (
	"fmt"
	"os"
)

// ThrowError to print the error message and exit the runtime
func ThrowError(value ...interface{}) {
	defer os.Exit(1)
	fmt.Println(value...)
}
