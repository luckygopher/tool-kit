// @Description: 工具箱
// @Author: Arvin
// @Date: 2021/3/8 4:26 下午
package main

import (
	"fmt"
	"os"

	rootCmd "github.com/qingyunjun/tool-kit/cmd/root"
)

func main() {
	// 工具入口
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
