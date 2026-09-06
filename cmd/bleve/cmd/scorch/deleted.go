package scorch

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deletedCmd = &cobra.Command{
	Use:   "deleted",
	Short: "deleted prints the deleted bitmap for segments in the index snapshot",
	Long:  `The delete command prints the deleted bitmap for segments in the index snapshot.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return fmt.Errorf("snapshot epoch required")
		} else if len(args) < 3 {
			snapshotEpoch, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			snapshot, err := index.LoadSnapshot(snapshotEpoch)
			if err != nil {
				return err
			}
			segments := snapshot.Segments()
			for i, segmentSnap := range segments {
				deleted := segmentSnap.Deleted()
				fmt.Printf("%d %v\n", i, deleted)
			}
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(deletedCmd)
}
