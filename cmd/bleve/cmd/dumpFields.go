package cmd

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/index/upsidedown"
	"github.com/spf13/cobra"
)

var dumpFieldsCmd = &cobra.Command{
	Use:   "fields [index path]",
	Short: "dump only the field rows",
	Long:  `The fields sub-command of dump will only dump the field rows.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := idx.Advanced()
		if err != nil {
			return fmt.Errorf("error getting index: %v", err)
		}
		r, err := i.Reader()
		if err != nil {
			return fmt.Errorf("error getting index reader: %v", err)
		}
		upsideDownReader, ok := r.(*upsidedown.IndexReader)
		if !ok {
			return fmt.Errorf("dump fields is only supported by index type upsidedown")
		}

		dumpChan := upsideDownReader.DumpFields()
		for rowOrErr := range dumpChan {
			switch rowOrErr := rowOrErr.(type) {
			case error:
				return fmt.Errorf("error dumping: %v", rowOrErr)
			case upsidedown.UpsideDownCouchRow:
				fmt.Printf("%v\n", rowOrErr)
				fmt.Printf("Key:   % -100x\nValue: % -100x\n\n", rowOrErr.Key(), rowOrErr.Value())
			}
		}
		return nil
	},
}

func init() {
	dumpCmd.AddCommand(dumpFieldsCmd)
}
