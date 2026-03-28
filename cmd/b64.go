package cmd

import (
	"fmt"

	"github.com/luckygopher/tool-kit/pkg/strutil"
	"github.com/urfave/cli/v2"
)

func b64Cmd() *cli.Command {
	return &cli.Command{
		Name:  "b64",
		Usage: "Base64 加解密",
		Subcommands: []*cli.Command{
			b64EncodeCmd(),
			b64DecodeCmd(),
		},
	}
}

func b64EncodeCmd() *cli.Command {
	return &cli.Command{
		Name:      "encode",
		Usage:     "Base64 编码",
		ArgsUsage: "<text>",
		Action: func(ctx *cli.Context) error {
			text := ctx.Args().First()
			if text == "" {
				return fmt.Errorf("text argument is required")
			}
			fmt.Println(strutil.B64Encode(text))
			return nil
		},
	}
}

func b64DecodeCmd() *cli.Command {
	return &cli.Command{
		Name:      "decode",
		Usage:     "Base64 解码",
		ArgsUsage: "<text>",
		Action: func(ctx *cli.Context) error {
			text := ctx.Args().First()
			if text == "" {
				return fmt.Errorf("text argument is required")
			}
			result, err := strutil.B64Decode(text)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, b64Cmd())
}
