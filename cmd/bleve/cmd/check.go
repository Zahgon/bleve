package cmd

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve/v2"
	"github.com/spf13/cobra"
)

var checkFieldName string
var checkCount int

var checkCmd = &cobra.Command{
	Use:   "check [index path]",
	Short: "checks the contents of the index",
	Long:  `The check command will perform consistency checks on the index.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var fieldNames []string
		var err error
		if checkFieldName == "" {
			fieldNames, err = idx.Fields()
			if err != nil {
				return err
			}
		} else {
			fieldNames = []string{checkFieldName}
		}
		fmt.Printf("checking fields: %v\n", fieldNames)

		totalProblems := 0
		for _, fieldName := range fieldNames {
			fmt.Printf("checking field: '%s'\n", fieldName)
			problems, err := checkField(idx, fieldName)
			if err != nil {
				log.Fatal(err)
			}
			totalProblems += problems
		}

		if totalProblems != 0 {
			return fmt.Errorf("found %d total problems\n", totalProblems)
		}

		return nil
	},
}

func checkField(index bleve.Index, fieldName string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getDictionary(index bleve.Index, field string) (map[string]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func init() {
	RootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringVarP(&checkFieldName, "field", "f", "", "Restrict check to the specified field name.")
	checkCmd.Flags().IntVarP(&checkCount, "count", "c", 100, "Check this many terms.")
}
