package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var keepDir, keepExt, parseJSON bool

var indexCmd = &cobra.Command{
	Use:   "index [index path] [data paths ...]",
	Short: "adds the files to the index",
	Long:  `The index command adds the specified files to the index.`,
	Annotations: map[string]string{
		canMutateBleveIndex: "true",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("must specify at least one path")
		}
		for file := range handleArgs(args[1:]) {
			var doc interface{}

			docID := file.filename
			if !keepDir {
				_, docID = filepath.Split(docID)
			}
			if !keepExt {
				ext := filepath.Ext(docID)
				docID = docID[0 : len(docID)-len(ext)]
			}
			doc = file.contents
			var err error
			if parseJSON {
				err = json.Unmarshal(file.contents, &doc)
				if err != nil {
					return fmt.Errorf("error parsing JSON: %v", err)
				}
			}
			fmt.Printf("Indexing: %s\n", docID)
			err = idx.Index(docID, doc)
			if err != nil {
				return fmt.Errorf("error indexing: %v", err)
			}
		}
		return nil
	},
}

type file struct {
	filename string
	contents []byte
}

func handleArgs(args []string) chan file { _ = "STUB: not implemented"; return nil }

func getAllFiles(args []string, rv chan file) { _ = "STUB: not implemented"; return }

func init() {
	RootCmd.AddCommand(indexCmd)

	indexCmd.Flags().BoolVarP(&keepDir, "keepDir", "d", false, "Keep the directory in the document id.")
	indexCmd.Flags().BoolVarP(&keepExt, "keepExt", "x", false, "Keep the extension in the document id.")
	indexCmd.Flags().BoolVarP(&parseJSON, "json", "j", true, "Parse the contents as JSON.")
}
