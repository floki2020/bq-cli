package cmd

import (
	"fmt"
	"github.com/floki2020/bq-cli/pkg"
	"github.com/floki2020/bq-cli/pkg/common"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show bq-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		tip := `bq-cli version:` + pkg.Green(common.Version)
		fmt.Println(tip)
	},
}
