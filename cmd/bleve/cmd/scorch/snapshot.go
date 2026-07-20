package scorch

import (
	"fmt"
	"strconv"

	seg "github.com/blevesearch/scorch_segment_api/v2"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "info prints details about the snapshots in the index",
	Long:  `The snapshot command prints details about the snapshots in the index.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			snapshotEpochs, err := index.RootBoltSnapshotEpochs()
			if err != nil {
				return err
			}
			for _, snapshotEpoch := range snapshotEpochs {
				fmt.Printf("snapshot epoch: %d\n", snapshotEpoch)
			}
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
				segment := segmentSnap.Segment()
				if segment, ok := segment.(seg.PersistedSegment); ok {
					fmt.Printf("%d %s\n", i, segment.Path())
				}
			}
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(snapshotCmd)
}
