package cmd

import (
	"github.com/blevesearch/bleve/v2/cmd/bleve/cmd/scorch"
)

func init() {
	RootCmd.AddCommand(scorch.RootCmd)
}
