//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/cmd/bleve/cmd"

	"github.com/spf13/cobra/doc"
)

func main() {
	cmd.RootCmd.DisableAutoGenTag = true
	identity := func(s string) string {
		return fmt.Sprintf(`{{< relref "docs/%s" >}}`, s)
	}
	emptyStr := func(s string) string { return "" }
	doc.GenMarkdownTreeCustom(cmd.RootCmd, "./", emptyStr, identity)
}
