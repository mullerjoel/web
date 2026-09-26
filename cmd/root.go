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
	rootCmd.Flags().BoolVarP(&openFlag, "open", "o", false, "open the url")
	rootCmd.Flags().BoolVarP(&copyFlag, "copy", "c", false, "copy the url to clipboard")
	rootCmd.Flags().BoolVarP(&printFlag, "print", "p", false, "print the url to stdout")
}
