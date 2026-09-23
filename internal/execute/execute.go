package execute

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"text/tabwriter"
	"web/internal/check"
	"web/internal/reader"
)

func renderLines(items []reader.Item) (lines []string, urlByLine map[string]string) {
	var buf bytes.Buffer
	format := tabwriter.NewWriter(&buf, 0, 0, 6, ' ', 0)
	for _, item := range items {
		fmt.Fprintf(format, "%s\t%s\t%s\n", item.Category, item.Name, item.Desc)
	}
	format.Flush()

	lines = strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")

	urlByLine = make(map[string]string, len(items))
	for i, line := range lines {
		urlByLine[line] = items[i].URL
	}

	return lines, urlByLine
}

func runFzf(lines []string) string {
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(lines, "\n"))

	out, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) && exitErr.ExitCode() == 130 {
			check.CheckError(errors.New("no selection made"))
		}
		check.CheckError(errors.New("fzf failed"))
	}

	return strings.TrimRight(string(out), "\n")
}

func RunFzf(items []reader.Item) string {
	lines, urlByLine := renderLines(items)

	selected := runFzf(lines)

	url, ok := urlByLine[selected]
	if !ok {
		check.CheckError(errors.New("no selection made"))
	}

	return url
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
