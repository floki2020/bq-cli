package cmd

import (
	"bytes"
	"fmt"
	"github.com/floki2020/bq-cli/pkg"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"os"
)

var (
	Name  string
	Dir   string
	Modal bool
)
var tableCmd = &cobra.Command{
	Use:   "view",
	Short: "create a basic vue3 page with table and form",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	tableCmd.PersistentFlags().StringVarP(&Name, "create", "c", "VuePage", "page name")
	tableCmd.PersistentFlags().StringVarP(&Dir, "dir", "d", "", "page path defaults views")
	tableCmd.PersistentFlags().BoolVarP(&Modal, "modal", "m", false, "without modal component")
}

func run() {
	//创建目录
	//创建文件

	pwd, err := os.Getwd()
	path := pwd + "/src/views/" + Dir + "/" + Name
	savePath := path + "/index.ts"

	if b := pkg.PathIsExits(path); !b {
		pkg.PathCreate(path)
	}

	cur, _ := pkg.GetExecutablePath()

	t := cur + "/.bqcli/tpl/vue3/types.ts.template"
	data, err := template.ParseFiles(t)
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	err = data.Execute(&b, Name)
	if err != nil {
		fmt.Print(err)
	}
	pkg.CreateFile(b, savePath)
}
