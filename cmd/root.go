package cmd

import (
	"github.com/spf13/cobra"
	"time"
)

var from string
var outputFormat string
var unixNano bool
var unix bool

const inputTimeFormat = "2006-01-02T15:04:05"
const outputTimeFormat = time.RFC3339

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "when",
	Short: "Time utility to get exact time in the past or future",
	Long: `Time utility that I'll describe when I figure out what it does'

It should be super cool and make my life easy`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "Base time for comparison in the layout "+inputTimeFormat)
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", outputTimeFormat, "Layout in which to output the time")
	rootCmd.PersistentFlags().BoolVarP(&unixNano, "unixnano", "n", false, "output time as UnixNano Timestamp")
	rootCmd.PersistentFlags().BoolVarP(&unix, "unix", "u", false, "output time as Unix Timestamp")
}
