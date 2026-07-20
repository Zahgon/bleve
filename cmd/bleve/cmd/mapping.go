package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var mappingCmd = &cobra.Command{
	Use:   "mapping [index path]",
	Short: "prints the mapping used for this index",
	Long:  `The mapping command prints a JSON representation of the mapping used for this index.`,
	Run: func(cmd *cobra.Command, args []string) {
		mapping := idx.Mapping()
		jsonBytes, err := json.MarshalIndent(mapping, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", jsonBytes)
	},
}

func init() {
	RootCmd.AddCommand(mappingCmd)
}
