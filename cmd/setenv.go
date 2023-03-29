package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

var setenvCmd = &cobra.Command{
	Use:   "setenv",
	Short: "set the cli execute path to environment",
	Run: func(cmd *cobra.Command, args []string) {
		runSetEnv()
	},
}

func runSetEnv() {
	p, _ := os.Executable()
	dir := filepath.Dir(p)

	sysPath := os.Getenv("PATH")

	newPath := dir + ";" + sysPath
	cmd := exec.Command("setx", "PATH", newPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

}
