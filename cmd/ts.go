package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

const timeLayout = "2006-01-02 15:04:05"

func tsCmd() *cli.Command {
	return &cli.Command{
		Name:  "ts",
		Usage: "时间戳工具",
		Subcommands: []*cli.Command{
			tsNowCmd(),
			tsToCmd(),
			tsFromCmd(),
		},
	}
}

func tsNowCmd() *cli.Command {
	return &cli.Command{
		Name:  "now",
		Usage: "显示当前时间（Unix 时间戳 + 多格式）",
		Action: func(ctx *cli.Context) error {
			now := time.Now()
			fmt.Printf("Unix:  %d\n", now.Unix())
			fmt.Printf("UnixMilli: %d\n", now.UnixMilli())
			fmt.Printf("UTC:   %s\n", now.UTC().Format(timeLayout))
			fmt.Printf("Local: %s\n", now.Local().Format(timeLayout))
			return nil
		},
	}
}

func tsToCmd() *cli.Command {
	return &cli.Command{
		Name:      "to",
		Usage:     "Unix 时间戳 → 可读时间",
		ArgsUsage: "<timestamp>",
		Action: func(ctx *cli.Context) error {
			arg := ctx.Args().First()
			if arg == "" {
				return fmt.Errorf("timestamp argument is required")
			}
			ts, err := strconv.ParseInt(arg, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid timestamp: %s", arg)
			}
			// Support both second and millisecond timestamps
			var t time.Time
			if ts > 1e12 {
				t = time.UnixMilli(ts)
			} else {
				t = time.Unix(ts, 0)
			}
			fmt.Printf("UTC:   %s\n", t.UTC().Format(timeLayout))
			fmt.Printf("Local: %s\n", t.Local().Format(timeLayout))
			return nil
		},
	}
}

func tsFromCmd() *cli.Command {
	return &cli.Command{
		Name:      "from",
		Usage:     "可读时间 → Unix 时间戳",
		ArgsUsage: "<datetime>",
		Action: func(ctx *cli.Context) error {
			arg := ctx.Args().First()
			if arg == "" {
				return fmt.Errorf("datetime argument is required (format: \"2006-01-02 15:04:05\")")
			}
			t, err := time.ParseInLocation(timeLayout, arg, time.Local)
			if err != nil {
				return fmt.Errorf("invalid datetime %q, expected format: 2006-01-02 15:04:05", arg)
			}
			fmt.Printf("Unix:      %d\n", t.Unix())
			fmt.Printf("UnixMilli: %d\n", t.UnixMilli())
			return nil
		},
	}
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, tsCmd())
}
