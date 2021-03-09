// @Description:
// @Author: Arvin
// @Date: 2021/3/9 11:46 下午
package convert_struct_cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// TODO 需要实现
var Cmd = &cobra.Command{
	Use:   "tts",
	Short: "database table convert struct",
	Long:  "database table convert struct",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(111)
		return nil
	},
}
