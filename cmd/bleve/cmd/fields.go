package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fieldsCmd = &cobra.Command{
	Use:   "fields [index path]",
	Short: "lists the fields in this index",
	Long:  `The fields command will list the fields used in this index.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		i, err := idx.Advanced()
		if err != nil {
			return fmt.Errorf("error getting index: %v", err)
		}
		r, err := i.Reader()
		if err != nil {
			return fmt.Errorf("error getting index reader: %v", err)
		}
		fields, err := r.Fields()
		if err != nil {
			return fmt.Errorf("error getting fields: %v", err)
		}
		for i, field := range fields {
			fmt.Printf("%d - %s\n", i, field)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(fieldsCmd)
}
