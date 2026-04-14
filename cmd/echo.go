package cmd

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

func echoCmd() *cli.Command {
	var (
		noNewline bool
		sep       string
	)
	return &cli.Command{
		Name:      "echo",
		Usage:     "输入内容并打印输出",
		ArgsUsage: "<text> [text...]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "no-newline",
				Aliases:     []string{"n"},
				Usage:       "不输出末尾换行符",
				Destination: &noNewline,
			},
			&cli.StringFlag{
				Name:        "sep",
				Aliases:     []string{"s"},
				Value:       " ",
				Usage:       "多个参数之间的分隔符",
				Destination: &sep,
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.Args().Len() == 0 {
				return fmt.Errorf("请提供要打印的内容，用法: echo <text> [text...]")
			}
			output := strings.Join(ctx.Args().Slice(), sep)
			if noNewline {
				fmt.Print(output)
			} else {
				fmt.Println(output)
			}
			return nil
		},
	}
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, echoCmd())
}
