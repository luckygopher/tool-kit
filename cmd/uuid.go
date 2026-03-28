package cmd

import (
	"fmt"

	"github.com/luckygopher/tool-kit/pkg/strutil"
	"github.com/urfave/cli/v2"
)

func uuidCmd() *cli.Command {
	var (
		count   int
		version string
		noDash  bool
		upper   bool
	)
	return &cli.Command{
		Name:  "uuid",
		Usage: "UUID 生成工具",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "n",
				Value:       1,
				Usage:       "生成数量",
				Destination: &count,
			},
			&cli.StringFlag{
				Name:        "version",
				Aliases:     []string{"v"},
				Value:       "v4",
				Usage:       "UUID 版本: v1, v4",
				Destination: &version,
			},
			&cli.BoolFlag{
				Name:        "no-dash",
				Usage:       "去掉连字符",
				Destination: &noDash,
			},
			&cli.BoolFlag{
				Name:        "upper",
				Aliases:     []string{"u"},
				Usage:       "大写输出",
				Destination: &upper,
			},
		},
		Action: func(ctx *cli.Context) error {
			if count < 1 {
				return fmt.Errorf("count must be >= 1")
			}
			for i := 0; i < count; i++ {
				var (
					id  string
					err error
				)
				switch version {
				case "v4", "4":
					id = strutil.UUID()
				case "v1", "1":
					id, err = strutil.UUIDV1()
					if err != nil {
						return err
					}
				default:
					return fmt.Errorf("unsupported version: %s (use v1 or v4)", version)
				}
				fmt.Println(strutil.FormatUUID(id, noDash, upper))
			}
			return nil
		},
	}
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, uuidCmd())
}
