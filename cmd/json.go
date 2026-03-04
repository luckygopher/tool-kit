package cmd

import (
	"fmt"

	"github.com/luckygopher/tool-kit/pkg/jsonutil"
	"github.com/urfave/cli/v2"
)

func jsonCmd() *cli.Command {
	return &cli.Command{
		Name:  "json",
		Usage: "JSON 工具",
		Subcommands: []*cli.Command{
			jsonFmtCmd(),
			jsonMiniCmd(),
			jsonValidCmd(),
		},
	}
}

func jsonFmtCmd() *cli.Command {
	var filePath string
	return &cli.Command{
		Name:  "fmt",
		Usage: "格式化 JSON（支持文件或 stdin）",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "input JSON file (default: stdin)",
				Destination: &filePath,
			},
		},
		Action: func(ctx *cli.Context) error {
			result, err := jsonutil.Format(filePath)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func jsonMiniCmd() *cli.Command {
	var filePath string
	return &cli.Command{
		Name:  "mini",
		Usage: "压缩 JSON",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "input JSON file (default: stdin)",
				Destination: &filePath,
			},
		},
		Action: func(ctx *cli.Context) error {
			result, err := jsonutil.Minify(filePath)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func jsonValidCmd() *cli.Command {
	var filePath string
	return &cli.Command{
		Name:  "valid",
		Usage: "校验 JSON 合法性",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "input JSON file (default: stdin)",
				Destination: &filePath,
			},
		},
		Action: func(ctx *cli.Context) error {
			if err := jsonutil.Validate(filePath); err != nil {
				return err
			}
			fmt.Println("valid JSON")
			return nil
		},
	}
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, jsonCmd())
}
