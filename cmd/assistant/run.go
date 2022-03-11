package main

import (
	"github.com/dafsic/assistant/config"
	"github.com/dafsic/assistant/lib/mylog"
	"github.com/dafsic/assistant/node"
	"github.com/dafsic/assistant/node/web"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
)

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "Auto pledge",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			EnvVars: []string{"LOTUS_ASSISTANT_CNF"},
			Value:   "~/.lotusassistant/config.toml",
			Usage:   "Load configuration from `FILE`",
		},
	},
	Action: func(cctx *cli.Context) error {
		// 要测试下多个模块都需要的情况下是不是初始化多个实例还是一个，文章中看是多个，只能有一个实例的话就需要sync.Once.
		// 目前看只会调用一次构造函数，不知道是不是因为返回值都是指针
		fx.New(
			fx.Supply(cctx), //config 模块需要从命令行参数中获取配置文件路径
			config.AssitantModule,
			mylog.Module,
			node.AssistantModule,
			web.RouterModule,
			web.HandlerModule,
			web.Register,
			fx.NopLogger,
		).Run() // 模块里不能有阻塞的协程，都要用go开启一个新的线程，Run()函数会在app.start后卡住等信号，收到中断信号会调用app.stop

		return nil
	},
}
