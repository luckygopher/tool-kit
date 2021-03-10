// @Description: 单词格式转换
// @Author: Arvin
// @Date: 2021/3/10 11:12 下午
package word

import "github.com/spf13/cobra"

var (
	str    string // 字符串
	format uint32 // 格式
)

var WordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  "单词格式转换",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
