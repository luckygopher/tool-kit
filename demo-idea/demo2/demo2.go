// @Description:
// @Author: Arvin
// @Date: 2021/3/9 4:34 下午
package demo2

import (
	"flag"
	"fmt"
	"os"

	"github.com/qingyunjun/tool-kit/demo-idea/command"
)

// 创建第一个子命令
var (
	age  int
	name string
)

var Demo = command.Command{
	Name: "demo2",
	InitCommand: func() error {
		fSet := flag.NewFlagSet("demo2", flag.ExitOnError)
		fSet.IntVar(&age, "age", 0, "年龄")
		fSet.StringVar(&name, "name", "", "姓名")
		return fSet.Parse(os.Args[2:])
	},
	Run: func() error {
		fmt.Println(age, name)
		return nil
	},
}
