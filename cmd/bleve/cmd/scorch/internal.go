package scorch

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var ascii bool

var internalCmd = &cobra.Command{
	Use:   "internal",
	Short: "internal prints the internal k/v pairs in a snapshot",
	Long:  `The internal command prints the internal k/v pairs in a snapshot.`,
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
			internal := snapshot.Internal()
			for k, v := range internal {
				if ascii {
					fmt.Printf("%s %s\n", k, string(v))
				} else {
					fmt.Printf("%x %x\n", k, v)
				}
			}
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(internalCmd)
	internalCmd.Flags().BoolVarP(&ascii, "ascii", "a", false, "print key/value in ascii")
}
