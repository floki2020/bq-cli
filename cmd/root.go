package cmd

import (
	"fmt"
	"github.com/floki2020/bq-cli/pkg"
	"github.com/floki2020/bq-cli/pkg/common"
	"github.com/spf13/cobra"
	"os"
)

//var projectBase string
//var usageTpl string

var rootCmd = &cobra.Command{
	Use:   "bq-cli",
	Short: "bq-cli is a tool for quickly building Vue 3 SPA .",
	Long:  "A cli tool to generate api,vue,ts  \n\n" + pkg.Green(common.Logo),
}

func init() {
	rootCmd.AddCommand(tableCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(setenvCmd)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
