// @Description:
// @Author: Arvin
// @Date: 2021/3/9 4:34 下午
package demo2

import (
	"flag"
	"fmt"
	"os"

	"github.com/qingyunjun/tool-kit/realization-idea/command"
)

// 创建第一个子命令
var age int

var Demo = command.Command{
	Name: "demo2",
	InitCommand: func() error {
		fSet := flag.NewFlagSet("age", flag.ExitOnError)
		fSet.IntVar(&age, "age", 0, "年龄")
		return fSet.Parse(os.Args[2:])
	},
	Run: func() error {
		fmt.Println(age)
		return nil
	},
}
