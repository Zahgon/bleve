package scorch

import (
	"fmt"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "info prints basic info about the index",
	Long:  `The info command prints basic info about the index.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		reader, err := index.Reader()
		if err != nil {
			return err
		}

		count, err := reader.DocCount()
		if err != nil {
			return err
		}

		fmt.Printf("doc count: %d\n", count)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)
}
