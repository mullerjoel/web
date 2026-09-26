package gitclone

import (
	"strings"
	"web/internal/shared"
)

type repoRef struct {
	Host string
	Path string
}

func FilterNoGit(items []shared.Item) []shared.Item {
	result := []shared.Item{}
	for _, item := range items {
		if item.Git != "" {
			result = append(result, item)
		}
	}
	return result
}

func GetItemRepo(items []shared.Item) []shared.Item {
	itemRepos := []shared.Item{}
	for _, item := range items {
		url := item.Git
		if url == "" {
			url = item.URL
		}
		ref := ParseRepoRef(url)
		if ref == nil {
			continue
		}
		item.GitHttps = ref.HTTPSCloneURL()
		item.Gitssh = ref.SSHCloneURL()
		itemRepos = append(itemRepos, item)
	}
	return itemRepos
}

func (r repoRef) SSHCloneURL() string {
	return "git@" + r.Host + ":" + r.Path + ".git"
}

func (r repoRef) HTTPSCloneURL() string {
	return "https://" + r.Host + "/" + r.Path + ".git"
}

func ParseRepoRef(rawURL string) *repoRef {
	rest := rawURL
	switch {
	case strings.HasPrefix(rest, "https://"):
		rest = strings.TrimPrefix(rest, "https://")
	case strings.HasPrefix(rest, "git@"):
		rest = strings.TrimPrefix(rest, "git@")
	default:
		return nil
	}
	rest = strings.TrimSuffix(rest, ".git")

	sep := strings.IndexAny(rest, ":/")
	if sep == -1 {
		return nil
	}

	return &repoRef{Host: rest[:sep], Path: rest[sep+1:]}
}
