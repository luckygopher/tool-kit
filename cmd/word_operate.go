package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func wordOperate1Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "action1",
		Usage: "test action",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			fmt.Println("test action1")
			return nil
		},
	}
	return cmd
}

func wordOperate2Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "action2",
		Usage: "test action",
		Flags: []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			fmt.Println("test action2")
			return nil
		},
	}
	return cmd
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, wordOperate1Cmd(), wordOperate2Cmd())
}
