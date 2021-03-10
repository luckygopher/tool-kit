// @Description:
// @Author: Arvin
// @Date: 2021/3/9 3:56 下午
package demo1

import (
	"flag"
	"fmt"
	"os"

	"github.com/qingyunjun/tool-kit/demo-idea/command"
)

// 创建第一个子命令
var name string

var Demo = command.Command{
	Name: "demo",
	InitCommand: func() error {
		fSet := flag.NewFlagSet("demo", flag.ExitOnError)
		fSet.StringVar(&name, "name", "aaa", "姓名")
		return fSet.Parse(os.Args[2:])
	},
	Run: func() error {
		fmt.Println(name)
		return nil
	},
}
