// @Description: 基础定义
// @Author: Arvin
// @Date: 2021/3/9 4:43 下午
package command

// 定义命令集
var Commands = make(map[string]Command, 0)

// 定义命令
type Command struct {
	Name        string
	InitCommand func() error
	Run         func() error
}
