package gitclone

import (
	"net/url"
	"strings"
	"web/internal/reader"
)

// Create a PR if there is one missing
var knownGitHosts = map[string]bool{
	"github.com":             true,
	"gitlab.com":             true,
	"bitbucket.org":          true,
	"codeberg.org":           true,
	"git.sr.ht":              true,
	"gitea.com":              true,
	"gitee.com":              true,
	"salsa.debian.org":       true,
	"gitlab.gnome.org":       true,
	"gitlab.freedesktop.org": true,
}

func FillRepos(items []reader.Item, ssh bool) []reader.Item {
	result := make([]reader.Item, 0, len(items))

	for _, item := range items {
		if item.Git != "" {
			result = append(result, item)
			continue
		}

		ref := ParseRepoRef(item.URL)
		if ref == nil {
			continue
		}

		if ssh {
			item.Git = ref.SSHCloneURL()
		} else {
			item.Git = ref.HTTPSCloneURL()
		}

		result = append(result, item)
	}

	return result
}

type RepoRef struct {
	Host string
	Path string
}

func (r RepoRef) SSHCloneURL() string {
	return "git@" + r.Host + ":" + r.Path + ".git"
}

func (r RepoRef) HTTPSCloneURL() string {
	return "https://" + r.Host + "/" + r.Path + ".git"
}

func ParseRepoRef(rawURL string) *RepoRef {
	u, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return nil
	}

	host := u.Hostname()
	if !knownGitHosts[host] {
		return nil
	}

	path, ok := normalizeRepoPath(u.Path)
	if !ok {
		return nil
	}

	return &RepoRef{Host: host, Path: path}
}

func normalizeRepoPath(rawPath string) (string, bool) {
	p := strings.Trim(rawPath, "/")
	p = strings.TrimSuffix(p, ".git")

	if p == "" {
		return "", false
	}

	if strings.Contains(p, "//") {
		return "", false
	}
	return p, true
}
