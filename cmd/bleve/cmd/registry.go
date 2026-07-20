package cmd

import (
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/spf13/cobra"
)

var registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "registry lists the bleve components compiled into this executable",
	Long:  `The registry command will list all of the bleve components compiled into this executable.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		types, instances := registry.CharFilterTypesAndInstances()
		printType("Char Filter", types, instances)

		types, instances = registry.TokenizerTypesAndInstances()
		printType("Tokenizer", types, instances)

		types, instances = registry.TokenMapTypesAndInstances()
		printType("Token Map", types, instances)

		types, instances = registry.TokenFilterTypesAndInstances()
		printType("Token Filter", types, instances)

		types, instances = registry.AnalyzerTypesAndInstances()
		printType("Analyzer", types, instances)

		types, instances = registry.DateTimeParserTypesAndInstances()
		printType("Date Time Parser", types, instances)

		types, instances = registry.KVStoreTypesAndInstances()
		printType("KV Store", types, instances)

		types, instances = registry.FragmentFormatterTypesAndInstances()
		printType("Fragment Formatter", types, instances)

		types, instances = registry.FragmenterTypesAndInstances()
		printType("Fragmenter", types, instances)

		types, instances = registry.HighlighterTypesAndInstances()
		printType("Highlighter", types, instances)
	},
}

func printType(label string, types, instances []string) { _ = "STUB: not implemented"; return }

func init() {
	RootCmd.AddCommand(registryCmd)
}
