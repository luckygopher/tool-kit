package cmd

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/qingyunjun/tool-kit/pkg/document"

	"github.com/urfave/cli/v2"

	"github.com/pkg/errors"
	"github.com/qingyunjun/tool-kit/config"
	"github.com/qingyunjun/tool-kit/pkg/db"
)

func dbToDocxCmd() *cli.Command {
	var (
		filePath   string // 结果输出文件路径
		configPath string // 配置文件
	)
	filePathFlag := &cli.StringFlag{
		Name: "filePath", Usage: "结果输出文件路径", Aliases: []string{"f"}, Destination: &filePath,
	}
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "配置文件路径", Aliases: []string{"c"}, Destination: &configPath,
	}

	cmd := &cli.Command{
		Name:  "docx",
		Usage: "导出表结构为docx",
		Flags: []cli.Flag{
			filePathFlag,
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			writer := document.NewWriter()
			// 获取数据库对象
			dbClient := db.NewClient(config.C.Database)
			if _, err := dbClient.ConnectDB(); err != nil {
				return errors.Errorf("数据库连接失败:%s", err)
			}
			// 查询所有表
			tables, err := dbClient.GetTables()
			if err != nil {
				return errors.Errorf("查询所有表错误:%s", err)
			}
			for _, tb := range tables {
				// 获取表结构
				tbDDL, err := dbClient.GetTableDDL(tb.Name)
				if err != nil {
					zap.L().Error("获取表结构失败", zap.String("tableName", tb.Name))
					return err
				}
				writer.WriterTable(tb.Name, tb.Comment, tbDDL)
			}
			if err := writer.Save(filePath); err != nil {
				fmt.Printf("写入失败 %s", err)
				return err
			}
			return nil
		},
	}
	return cmd
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, dbToDocxCmd())
}
