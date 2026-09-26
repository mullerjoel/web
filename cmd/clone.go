/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"web/internal/execute"
	"web/internal/find"
	"web/internal/gitclone"
	"web/internal/reader"
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
	autoSsh        bool
	autoHttps      bool
	copyCloneLink  bool
	printCloneLink bool
)

func runClone(cmd *cobra.Command, args []string) {
	items := reader.Read()
	if autoSsh || autoHttps {
		items = gitclone.GetItemRepo(items)
	} 

	item := find.FindItem(items)
	url := item.Git

	if autoSsh {
		url = item.Gitssh
	}
	if autoHttps {
		url = item.GitHttps
	}

	// TODO: Handle the errors
	execute.CloneRepo(url)

	if copyCloneLink {
		execute.CopyUrl(url)
	}
	if printCloneLink {
		execute.PrintUrl(url)
	}
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().BoolVar(&autoSsh, "auto-ssh", false, "automatically detect repositories and clone the repository with ssh")
	cloneCmd.Flags().BoolVar(&autoHttps, "auto-https", false, "automatically detect repositories and clone the repository with https")
	cloneCmd.Flags().BoolVarP(&copyCloneLink, "copy", "c", false, "copy the clone url to clipboard")
	cloneCmd.Flags().BoolVarP(&printCloneLink, "print", "p", false, "print the clone url to stdout")

	cloneCmd.MarkFlagsMutuallyExclusive("auto-ssh", "auto-https")
}
