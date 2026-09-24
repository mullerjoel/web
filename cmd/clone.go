/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"web/internal/execute"
	"web/internal/gitclone"
	"web/internal/reader"

	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runClone,
}

var (
	sshFlag            bool
	httpsFlag          bool
	autoFlag           bool
	copyCloneLinkFlag  bool
	printCloneLinkFlag bool
)

func runClone(cmd *cobra.Command, args []string) {
	items := reader.Read()
	filtered := gitclone.FillRepos(items, sshFlag)
	url := execute.FindUrl(filtered, execute.Repo)

	execute.CloneRepo(url)

	if copyCloneLinkFlag {
		execute.CopyUrl(url)
	}
	if printCloneLinkFlag {
		execute.PrintUrl(url)
	}
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().BoolVar(&sshFlag, "ssh", false, "clone the repository with ssh")
	cloneCmd.Flags().BoolVar(&httpsFlag, "https", false, "clone the repository with https")
	cloneCmd.Flags().BoolVarP(&autoFlag, "auto", "a", false, "automatically detect repositories")
	cloneCmd.Flags().BoolVarP(&copyCloneLinkFlag, "copy", "c", false, "copy the clone url to clipboard")
	cloneCmd.Flags().BoolVarP(&printCloneLinkFlag, "print", "p", false, "print the clone url to stdout")

	cloneCmd.MarkFlagsMutuallyExclusive("ssh", "https")
}
