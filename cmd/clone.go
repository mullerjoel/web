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

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Fuzzy-find and git clone a bookmarked repository",
	Long: `Fuzzy-find a bookmark and git clone it.

By default only bookmarks with a "git" field are shown, cloned as-is.

--auto-ssh and --auto-https also include bookmarks with just a "url"
field, and rebuild the URL as SSH or HTTPS before cloning.`,
	Run: runClone,
}

var (
	autoSsh   bool
	autoHttps bool
)

func runClone(cmd *cobra.Command, args []string) {
	items := reader.Read()
	if autoSsh || autoHttps {
		items = gitclone.GetItemRepo(items)
	} else {
		items = gitclone.FilterNoGit(items)
	}

	item := find.FindItem(items)
	url := item.Git

	if autoSsh {
		url = item.Gitssh
	}
	if autoHttps {
		url = item.GitHttps
	}

	execute.CloneRepo(url)
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().BoolVar(&autoSsh, "auto-ssh", false, "use git field (or url as fallback) and clone via ssh")
	cloneCmd.Flags().BoolVar(&autoHttps, "auto-https", false, "use git field (or url as fallback) and clone via https")
	cloneCmd.MarkFlagsMutuallyExclusive("auto-ssh", "auto-https")
}
