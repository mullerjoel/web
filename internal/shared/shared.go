package shared

import (
	"errors"
	"fmt"
	"os"
)

type Item struct {
	Category string
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Desc     string `yaml:"desc,omitempty"`
	Git      string `yaml:"git,omitempty"`
	Gitssh   string
	GitHttps string
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func ThrowErrorWithMessage(message string) {
	fmt.Fprintln(os.Stderr, "error:", errors.New(message))
	os.Exit(1)
}

func ThrowError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
