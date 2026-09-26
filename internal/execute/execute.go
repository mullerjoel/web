package execute

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
	"web/internal/shared"
)

func RunFzf(lines []string) int {
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(lines, "\n"))

	out, err := cmd.Output()
	if err != nil {
		msg := "fzf failed"
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 130 {
			msg = "no selection made"
		}
		shared.ThrowErrorWithMessage(msg)
		return -1
	}

	return slices.Index(lines, strings.TrimRight(string(out), "\n"))
}

func OpenUrl(url string) error {
	var cmd string
	var args = []string{url}

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	return exec.Command(cmd, args...).Start()
}

func CloneRepo(url string) error {
	args := []string{"clone", url}
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CopyUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "pbcopy"
	case "windows":
		cmd = "clip"
	default: // linux
		cmd = "xclip"
		args = []string{"-selection", "clipboard"}
	}

	command := exec.Command(cmd, args...)
	command.Stdin = strings.NewReader(url)
	return command.Run()
}

func PrintUrl(url string) {
	fmt.Println(url)
}
