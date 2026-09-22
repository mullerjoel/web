package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "web",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ArbitraryArgs,
	Run:  runOpen,
}

var (
	openFlag  bool
	copyFlag  bool
	printFlag bool
)

func runFzf(items []string) (string, error) {
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(items, "\n"))
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) && exitErr.ExitCode() == 130 {
		return "", errors.New("no selection made")
	}
	if err != nil {
		return "", fmt.Errorf("fzf failed: %w", err)
	}

	result := strings.TrimSpace(string(out))
	if result == "" {
		return "", errors.New("no selection made")
	}
	return result, nil
}

func openUrl(url string) error {
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

func copyUrl(url string) error {
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

func printUrl(url string) {
	fmt.Println(url)
}

func runOpen(cmd *cobra.Command, args []string) {
	if !openFlag && !copyFlag && !printFlag {
		openFlag = true
	}

	urls := []string{"https://apple.com", "https://youtube.com", "https://google.com"}
	url, err := runFzf(urls)

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	if openFlag {
		openUrl(url)
	}
	if copyFlag {
		copyUrl(url)
	}
	if printFlag {
		printUrl(url)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&openFlag, "open", "o", false, "open the url")
	rootCmd.Flags().BoolVarP(&copyFlag, "copy", "c", false, "copy the url to clipboard")
	rootCmd.Flags().BoolVarP(&printFlag, "print", "p", false, "print the url to stdout")
}
