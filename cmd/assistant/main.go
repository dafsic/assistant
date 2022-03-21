package main

import (
	"github.com/dafsic/assistant/utils"
	"github.com/dafsic/assistant/version"
	"github.com/urfave/cli/v2"
	"os"
)

// 这里mylog还没有初始化，不能使用
var logger = utils.NewLogger(os.Stdout, "main", utils.LDebug, utils.Ldefault)

func main() {
	app := &cli.App{
		Name:    "lotus assistant",
		Usage:   "lotus-assistant command [args]",
		Version: version.AssistantVersion.String(),
		Commands: []*cli.Command{
			runCmd,
			pledgeCmd,
			switchCmd,
			//...
		},
	}
	if err := app.Run(os.Args); err != nil {
		logger.Warnf("%+v", err)
		return
	}
}
