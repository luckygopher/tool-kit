package cmd

import (
	"fmt"

	"github.com/luckygopher/tool-kit/pkg/strutil"
	"github.com/urfave/cli/v2"
)

func strCmd() *cli.Command {
	return &cli.Command{
		Name:  "str",
		Usage: "字符串工具箱",
		Subcommands: []*cli.Command{
			strHashCmd(),
			strB64Cmd(),
			strURLCmd(),
			strUUIDCmd(),
			strRandCmd(),
			strCaseCmd(),
		},
	}
}

func strHashCmd() *cli.Command {
	var hashType string
	return &cli.Command{
		Name:      "hash",
		Usage:     "哈希计算 (md5|sha256|sha1)",
		ArgsUsage: "<text>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "type",
				Aliases:     []string{"t"},
				Value:       "md5",
				Usage:       "hash type: md5, sha256, sha1",
				Destination: &hashType,
			},
		},
		Action: func(ctx *cli.Context) error {
			text := ctx.Args().First()
			if text == "" {
				return fmt.Errorf("text argument is required")
			}
			result, err := strutil.Hash(hashType, text)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func strB64Cmd() *cli.Command {
	return &cli.Command{
		Name:      "b64",
		Usage:     "Base64 编解码 (encode|decode)",
		ArgsUsage: "<encode|decode> <text>",
		Action: func(ctx *cli.Context) error {
			if ctx.Args().Len() < 2 {
				return fmt.Errorf("usage: str b64 <encode|decode> <text>")
			}
			op := ctx.Args().Get(0)
			text := ctx.Args().Get(1)
			switch op {
			case "encode":
				fmt.Println(strutil.B64Encode(text))
			case "decode":
				result, err := strutil.B64Decode(text)
				if err != nil {
					return err
				}
				fmt.Println(result)
			default:
				return fmt.Errorf("unknown operation: %s (use encode or decode)", op)
			}
			return nil
		},
	}
}

func strURLCmd() *cli.Command {
	return &cli.Command{
		Name:      "url",
		Usage:     "URL 编解码 (encode|decode)",
		ArgsUsage: "<encode|decode> <text>",
		Action: func(ctx *cli.Context) error {
			if ctx.Args().Len() < 2 {
				return fmt.Errorf("usage: str url <encode|decode> <text>")
			}
			op := ctx.Args().Get(0)
			text := ctx.Args().Get(1)
			switch op {
			case "encode":
				fmt.Println(strutil.URLEncode(text))
			case "decode":
				result, err := strutil.URLDecode(text)
				if err != nil {
					return err
				}
				fmt.Println(result)
			default:
				return fmt.Errorf("unknown operation: %s (use encode or decode)", op)
			}
			return nil
		},
	}
}

func strUUIDCmd() *cli.Command {
	return &cli.Command{
		Name:  "uuid",
		Usage: "生成 UUID v4",
		Action: func(ctx *cli.Context) error {
			fmt.Println(strutil.UUID())
			return nil
		},
	}
}

func strRandCmd() *cli.Command {
	var n int
	var typ string
	return &cli.Command{
		Name:  "rand",
		Usage: "生成随机字符串",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "n",
				Value:       32,
				Usage:       "string length",
				Destination: &n,
			},
			&cli.StringFlag{
				Name:        "type",
				Aliases:     []string{"t"},
				Value:       "mix",
				Usage:       "character set: alpha, num, mix",
				Destination: &typ,
			},
		},
		Action: func(ctx *cli.Context) error {
			result, err := strutil.RandString(n, typ)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func strCaseCmd() *cli.Command {
	return &cli.Command{
		Name:      "case",
		Usage:     "命名风格转换 (snake|camel|pascal)",
		ArgsUsage: "<snake|camel|pascal> <text>",
		Action: func(ctx *cli.Context) error {
			if ctx.Args().Len() < 2 {
				return fmt.Errorf("usage: str case <snake|camel|pascal> <text>")
			}
			style := ctx.Args().Get(0)
			text := ctx.Args().Get(1)
			var result string
			switch style {
			case "snake":
				result = strutil.ToSnake(text)
			case "camel":
				result = strutil.ToCamel(text)
			case "pascal":
				result = strutil.ToPascal(text)
			default:
				return fmt.Errorf("unknown style: %s (use snake, camel, pascal)", style)
			}
			fmt.Println(result)
			return nil
		},
	}
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, strCmd())
}
