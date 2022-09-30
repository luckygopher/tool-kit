package cmd

import (
	"fmt"

	"github.com/luckygopher/tool-kit/config"
	vaccineImp "github.com/luckygopher/tool-kit/pkg/vaccine"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func vaccine() *cli.Command {
	cmd := &cli.Command{
		Name:  "vaccine",
		Usage: "疫苗工具",
	}
	cmd.Subcommands = append(cmd.Subcommands, seckill())
	cmd.Subcommands = append(cmd.Subcommands, list())
	cmd.Subcommands = append(cmd.Subcommands, area())
	cmd.Subcommands = append(cmd.Subcommands, member())
	cmd.Subcommands = append(cmd.Subcommands, FindAreaVaccine())
	return cmd
}

func seckill() *cli.Command {
	var configPath string
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "指定配置文件", Aliases: []string{"c"}, Destination: &configPath,
	}

	cmd := &cli.Command{
		Name:  "seckill",
		Usage: "在秒苗小程序秒杀疫苗",
		Flags: []cli.Flag{
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			setups := Setup{
				SetLogger,
				SetupHTTP,
			}
			setups.apply()
			vaccineClient := vaccineImp.NewClient(config.C.Vaccine, zap.L())
			if err := vaccineClient.Start(); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func list() *cli.Command {
	var configPath string
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "指定配置文件", Aliases: []string{"c"}, Destination: &configPath,
	}

	cmd := &cli.Command{
		Name:  "list",
		Usage: "获取秒苗小程序某地区的疫苗列表",
		Flags: []cli.Flag{
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			setups := Setup{
				SetLogger,
				SetupHTTP,
			}
			setups.apply()
			vaccineClient := vaccineImp.NewClient(config.C.Vaccine, zap.L())
			res, err := vaccineClient.GetVaccineList(config.C.Vaccine.RegionCode)
			if err != nil {
				return err
			}
			for _, datum := range res.Data {
				fmt.Printf("ID:%d\n", datum.ID)
				fmt.Printf("Name:%s\n", datum.Name)
				fmt.Printf("Address:%s\n", datum.Address)
				fmt.Printf("VaccineCode:%s\n", datum.VaccineCode)
				fmt.Printf("VaccineName:%s\n", datum.VaccineName)
				fmt.Printf("StartTime:%s\n", datum.StartTime)
				fmt.Println("--------------")
			}
			return nil
		},
	}
	return cmd
}

func area() *cli.Command {
	var configPath string
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "指定配置文件", Aliases: []string{"c"}, Destination: &configPath,
	}

	cmd := &cli.Command{
		Name:  "area",
		Usage: "获取某省区域的所有市信息",
		Flags: []cli.Flag{
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			setups := Setup{
				SetLogger,
				SetupHTTP,
			}
			setups.apply()
			vaccineClient := vaccineImp.NewClient(config.C.Vaccine, zap.L())
			res, err := vaccineClient.GetArea(config.C.Vaccine.ParentCode)
			if err != nil {
				return err
			}
			for _, datum := range res.Data {
				fmt.Printf("%+v\n", datum)
			}
			return nil
		},
	}
	return cmd
}

func member() *cli.Command {
	var configPath string
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "指定配置文件", Aliases: []string{"c"}, Destination: &configPath,
	}

	cmd := &cli.Command{
		Name:  "member",
		Usage: "获取账户信息",
		Flags: []cli.Flag{
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			setups := Setup{
				SetLogger,
				SetupHTTP,
			}
			setups.apply()
			vaccineClient := vaccineImp.NewClient(config.C.Vaccine, zap.L())
			res, err := vaccineClient.GetMember()
			if err != nil {
				return err
			}
			for _, datum := range res.Data {
				fmt.Printf("%+v\n", datum)
			}
			return nil
		},
	}
	return cmd
}

func FindAreaVaccine() *cli.Command {
	var configPath string
	configPathFlag := &cli.StringFlag{
		Name: "configPath", Usage: "指定配置文件路径", Aliases: []string{"c"}, Destination: &configPath,
	}
	cmd := &cli.Command{
		Name:  "find",
		Usage: "查询某省当前有苗的地区及该地区可秒疫苗",
		Flags: []cli.Flag{
			configPathFlag,
		},
		Action: func(ctx *cli.Context) error {
			config.ParseConfig(configPath)
			setups := Setup{
				SetLogger,
				SetupHTTP,
			}
			setups.apply()
			vaccineClient := vaccineImp.NewClient(config.C.Vaccine, zap.L())
			if err := vaccineClient.FindAreaVaccine(config.C.Vaccine.ParentCode); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func init() {
	rootCmd.Commands = append(rootCmd.Commands, vaccine())
}
