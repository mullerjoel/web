package find

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"web/internal/execute"
	"web/internal/shared"
)

func renderLines(items []shared.Item) []string {
	if len(items) == 0 {
		return nil
	}

	var buf bytes.Buffer
	tw := tabwriter.NewWriter(&buf, 0, 0, 3, ' ', 0)
	for _, item := range items {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", item.Category, item.Name, item.Desc)
	}
	tw.Flush()

	return strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
}

func FindItem(items []shared.Item) shared.Item {
	lines := renderLines(items)
	if len(lines) < 1 {
		shared.ThrowErrorWithMessage("no urls read")
	}
	idx := execute.RunFzf(lines)
	if idx < 0 {
		shared.ThrowErrorWithMessage("fzf failed")
	}
	return items[idx]
}
