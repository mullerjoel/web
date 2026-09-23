package check

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
