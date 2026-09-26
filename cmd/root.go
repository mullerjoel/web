package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"web/internal/execute"
	"web/internal/find"
	"web/internal/reader"
)

var rootCmd = &cobra.Command{
	Use:   "web",
	Short: "Fuzzy-find and open bookmarked URLs",
	Long: `Fuzzy-find URLs bookmarked in ~/.config/web/*.yaml and open, copy, or print the one you pick.

Use "web clone" to clone one of your bookmarks as a git repo instead.`,
	Args: cobra.ArbitraryArgs,
	Run:  runOpen,
}

var (
	openFlag  bool
	copyFlag  bool
	printFlag bool
)

func runOpen(cmd *cobra.Command, args []string) {
	if !openFlag && !copyFlag && !printFlag {
		openFlag = true
	}

	items := reader.Read()
	item := find.FindItem(items)
	url := item.URL

	if openFlag {
		execute.OpenUrl(url)
	}
	if copyFlag {
		execute.CopyUrl(url)
	}
	if printFlag {
		execute.PrintUrl(url)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&openFlag, "open", "o", false, "open the selected url in your browser")
	rootCmd.Flags().BoolVarP(&copyFlag, "copy", "c", false, "copy the selected url to the clipboard")
	rootCmd.Flags().BoolVarP(&printFlag, "print", "p", false, "print the selected url to stdout")
}
