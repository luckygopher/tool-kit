// @Description: 多个子命令基础实现思路,作为学习笔记
// @Author: Arvin
// @Date: 2021/3/9 1:04 下午
package realization_idea

import (
	"os"

	"github.com/qingyunjun/tool-kit/realization-idea/demo2"

	"github.com/qingyunjun/tool-kit/realization-idea/command"

	"github.com/qingyunjun/tool-kit/realization-idea/demo1"
)

func init() {
	command.Commands["demo"] = demo1.Demo
	command.Commands["demo2"] = demo2.Demo
}

// 执行命令
func Run() {
	if len(os.Args) > 1 {
		c := command.Commands[os.Args[1]]
		c.InitCommand()
		c.Run()
	}
}
