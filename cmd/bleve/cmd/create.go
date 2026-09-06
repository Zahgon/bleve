package cmd

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/spf13/cobra"
)

var mappingPath, indexType, storeType string

var createCmd = &cobra.Command{
	Use:   "create [index path]",
	Short: "creates a new index",
	Long:  `The create command will create a new empty index.`,
	Annotations: map[string]string{
		canMutateBleveIndex: "true",
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 1 {
			return fmt.Errorf("must specify path to index")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var mapping mapping.IndexMapping
		var err error
		mapping, err = buildMapping()
		if err != nil {
			return fmt.Errorf("error building mapping: %v", err)
		}
		idx, err = bleve.NewUsing(args[0], mapping, indexType, storeType, nil)
		if err != nil {
			return fmt.Errorf("error creating index: %v", err)
		}

		return nil
	},
}

func buildMapping() (mapping.IndexMapping, error) {
	_ = "STUB: not implemented"
	return *new(mapping.IndexMapping), nil
}

func init() {
	RootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&mappingPath, "mapping", "m", "", "Path to a file containing a JSON representation of an index mapping to use.")
	createCmd.Flags().StringVarP(&storeType, "store", "s", bleve.Config.DefaultKVStore, "The bleve storage type to use.")
	createCmd.Flags().StringVarP(&indexType, "index", "i", bleve.Config.DefaultIndexType, "The bleve index type to use.")
}
