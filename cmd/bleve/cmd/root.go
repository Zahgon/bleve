package cmd

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/spf13/cobra"
)

var cfgFile string

var idx bleve.Index

var DefaultOpenReadOnly = false

const canMutateBleveIndex = "canMutateBleveIndex"

func CanMutateBleveIndex(c *cobra.Command) bool { _ = "STUB: not implemented"; return false }

var RootCmd = &cobra.Command{
	Use:   "bleve",
	Short: "command-line tool to interact with a bleve index",
	Long:  `Bleve is a command-line tool to interact with a bleve index.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Use == "bash" || cmd.Use == "zsh" || cmd.Use == "fish" || cmd.Use == "powershell" {

			return nil
		}
		if len(args) < 1 {
			return fmt.Errorf("must specify path to index")
		}
		runtimeConfig := map[string]interface{}{
			"read_only": DefaultOpenReadOnly,
		}
		var err error
		idx, err = bleve.OpenUsing(args[0], runtimeConfig)
		if err != nil {
			return fmt.Errorf("error opening bleve index: %v", err)
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Use == "bash" || cmd.Use == "zsh" || cmd.Use == "fish" || cmd.Use == "powershell" {

			return nil
		}
		err := idx.Close()
		if err != nil {
			return fmt.Errorf("error closing bleve index: %v", err)
		}
		return nil
	},
}

func Execute() { _ = "STUB: not implemented"; return }
