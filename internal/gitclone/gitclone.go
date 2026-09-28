package gitclone

import (
	"strings"
	"web/internal/shared"
)

type repoRef struct {
	Host string
	Path string
	Port string
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
		item = setGitURLs(*ref, item)
		itemRepos = append(itemRepos, item)
	}
	return itemRepos
}

func setGitURLs(ref repoRef, item shared.Item) shared.Item {
	if strings.HasPrefix(item.Git, "https://") {
		item.GitHttps = item.Git
	} else {
		item.GitHttps = ref.HTTPSCloneURL()
	}
	if isSSHURL(item.Git) {
		item.Gitssh = item.Git
	} else {
		item.Gitssh = ref.SSHCloneURL()
	}
	return item
}

func isSSHURL(url string) bool {
	if strings.HasPrefix(url, "ssh://") {
		return true
	}
	return !strings.Contains(url, "://") && strings.Contains(url, ":")
}

func (r repoRef) SSHCloneURL() string {
	if r.Port != "" {
		return "ssh://git@" + r.Host + ":" + r.Port + "/" + r.Path + ".git"
	}
	return "git@" + r.Host + ":" + r.Path + ".git"
}

func (r repoRef) HTTPSCloneURL() string {
	if r.Port != "" {
		return "https://" + r.Host + ":" + r.Port + "/" + r.Path + ".git"
	}
	return "https://" + r.Host + "/" + r.Path + ".git"
}

func splitPort(rest string) (string, string) {
	host, path, _ := strings.Cut(rest, "/")
	h, port, ok := strings.Cut(host, ":")
	if !ok {
		return rest, ""
	}
	return h + "/" + path, port
}

func ParseRepoRef(rawURL string) *repoRef {
	rest := rawURL
	port := ""
	switch {
	case strings.HasPrefix(rest, "https://"):
		rest = strings.TrimPrefix(rest, "https://")
		rest, port = splitPort(rest)
	case strings.HasPrefix(rest, "git@"):
		rest = strings.TrimPrefix(rest, "git@")
	case strings.HasPrefix(rest, "ssh://"):
		rest = strings.TrimPrefix(rest, "ssh://")
		rest = strings.TrimPrefix(rest, "git@")
		rest, port = splitPort(rest)
	default:
		return nil
	}
	rest = strings.TrimSuffix(rest, ".git")

	sep := strings.IndexAny(rest, ":/")
	if sep == -1 {
		return nil
	}

	return &repoRef{Host: rest[:sep], Path: rest[sep+1:], Port: port}
}
