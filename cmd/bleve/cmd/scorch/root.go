package scorch

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/index/scorch"
	"github.com/spf13/cobra"
)

var index *scorch.Scorch

var RootCmd = &cobra.Command{
	Use:   "scorch",
	Short: "command-line tool to interact with a scorch index",
	Long:  `Scorch is a command-line tool to interact with a scorch index.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return fmt.Errorf("must specify path to scorch index")
		}

		readOnly := true
		config := map[string]interface{}{
			"read_only": readOnly,
			"path":      args[0],
		}

		idx, err := scorch.NewScorch(scorch.Name, config, nil)
		if err != nil {
			return err
		}

		err = idx.Open()
		if err != nil {
			return fmt.Errorf("error opening: %v", err)
		}

		index = idx.(*scorch.Scorch)

		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() { _ = "STUB: not implemented"; return }
