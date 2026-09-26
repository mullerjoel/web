# web - open and clone your bookmarks fast

Fuzzy-find URLs bookmarked in `~/.config/web/*.yaml` and open, copy, or print the one you pick. Use `web clone` to clone one of your bookmarks as a Git repository instead.

## Usage

### Open, Print, and Copy URL

```txt
Usage:
  web [flags]

Flags:
  -c, --copy    copy the selected url to the clipboard
  -o, --open    open the selected url in your browser
  -p, --print   print the selected url to stdout
  -h, --help    help for web
```

### Clone Repository

```txt
Usage:
  web clone [flags]

Flags:
      --auto-https   use git field (or url as fallback) and clone via https
      --auto-ssh     use git field (or url as fallback) and clone via ssh
  -h, --help         help for clone
```

## Installation

### Homebrew

If you use Homebrew, you can install the app via the official cask:

```bash
brew install --cask mullerjoel/tap/web
```


### Other Platforms

For all other devices, download the latest version from the [web releases page](https://github.com/mullerjoel/web/releases).

### Build from Source

Make sure [Go](https://go.dev) is installed, then clone the repository and build the binary:

```bash
go build
```

You can then move the `web` binary somewhere in your `PATH`.

## Config File

Store multiple files in `~/.config/web/*.yaml`. The files should be structured as follows:

```yaml
repositories: # multiple labels possible, naming doesn't matter
  - name: web
    url: https://github.com/mullerjoel/web
    desc: a very nice binary # optional
    git: git@github.com:mullerjoel/web.git # optional
  - name: # ... multiple items possible
```

## Shell Completions

When the binary is installed with Homebrew, the completions are already installed. With another installation, the completion can be installed as follows. For more information, see the [Shell Completion Guide](https://cobra.dev/docs/how-to-guides/shell-completion/).

```bash
web completion bash
web completion zsh
web completion fish
web completion powershell
```
