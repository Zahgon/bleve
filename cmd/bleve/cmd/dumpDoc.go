package cmd

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/index/upsidedown"
	"github.com/spf13/cobra"
)

var dumpDocCmd = &cobra.Command{
	Use:   "doc [index path] [doc id]",
	Short: "dump only the rows relating to this doc ID",
	Long:  `The doc sub-command of dump will only dump the rows relating to this doc ID.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("must specify docid")
		}

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
			return fmt.Errorf("dump doc is only supported by index type upsidedown")
		}

		dumpChan := upsideDownReader.DumpDoc(args[1])
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
	dumpCmd.AddCommand(dumpDocCmd)
}
