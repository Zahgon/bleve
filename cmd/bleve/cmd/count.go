package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count [index path]",
	Short: "counts the number documents in the index",
	Long:  `The count command will count the number of documents in the index.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		count, err := idx.DocCount()
		if err != nil {
			return fmt.Errorf("error counting docs in index: %v", err)
		}
		fmt.Printf("%d\n", count)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(countCmd)
}
