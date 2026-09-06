package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dictionaryCmd = &cobra.Command{
	Use:   "dictionary [index path] [field name]",
	Short: "prints the term dictionary for the specified field in the index",
	Long:  `The dictionary command will print the term dictionary for the specified field.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("must specify field")
		}
		i, err := idx.Advanced()
		if err != nil {
			return fmt.Errorf("error getting index: %v", err)
		}
		r, err := i.Reader()
		if err != nil {
			return fmt.Errorf("error getting index reader: %v", err)
		}
		d, err := r.FieldDict(args[1])
		if err != nil {
			return fmt.Errorf("error getting field dictionary: %v", err)
		}

		de, err := d.Next()
		for err == nil && de != nil {
			fmt.Printf("%s - %d\n", de.Term, de.Count)
			de, err = d.Next()
		}
		if err != nil {
			return fmt.Errorf("error iterating dictionary: %v", err)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(dictionaryCmd)
}
