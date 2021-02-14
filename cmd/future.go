package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// futureCmd represents the future command
var futureCmd = &cobra.Command{
	Use:   "future",
	Short: "Time in the future",
	RunE: func(cmd *cobra.Command, args []string) error {
		fromTime, err := parseFromTime(from)
		if err != nil {
			return err
		}

		duration, err := time.ParseDuration(args[0])
		if err != nil {
			panic(err)
		}

		agoTime := fromTime.Add(duration)

		fmt.Println(outputResult(agoTime))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(futureCmd)
}
