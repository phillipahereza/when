package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// agoCmd represents the ago command
var agoCmd = &cobra.Command{
	Use:   "ago",
	Short: "Time in the past",
	RunE: func(cmd *cobra.Command, args []string) error {
		fromTime, err := parseFromTime(from)
		if err != nil {
			return err
		}

		duration, err := time.ParseDuration(args[0])
		if err != nil {
			panic(err)
		}

		agoTime := fromTime.Add(-duration)

		fmt.Println(outputResult(agoTime))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(agoCmd)
}

func parseFromTime(f string) (time.Time, error) {
	if f == "" {
		return time.Now(), nil
	}
	return time.Parse(inputTimeFormat, from)
}

func outputResult(t time.Time) interface{} {
	if unixNano {
		return t.UnixNano()
	} else if unix {
		return t.Unix()
	}
	return t.Format(time.RFC3339)
}
