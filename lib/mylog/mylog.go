// 日志的level等信息应该从配置文件里读取，配置文件的路径应该从环境变量或者命令行参数里获取
package mylog

import (
	"fmt"
	"github.com/dafsic/assistant/config"
	"github.com/dafsic/assistant/utils"
	"go.uber.org/fx"
	"io"
	"os"
	"sync"
)

type LoggingI interface {
	GetLogger(name string) *utils.Logger
}

type LoggingT struct {
	mux     sync.Mutex
	lvl     string
	output  io.Writer
	loggers map[string]*utils.Logger
}

func (l *LoggingT) GetLogger(name string) *utils.Logger {
	l.mux.Lock()
	defer l.mux.Unlock()
	i, ok := l.loggers[name]
	if !ok {
		i = utils.NewLogger(l.output, name, utils.LogLevelFromString(l.lvl), utils.Ldefault)
		l.loggers[name] = i
	}
	return i
}

var once sync.Once
var l LoggingI

func NewMylog(cfg config.ConfigI) LoggingI {
	once.Do(func() {
		fmt.Println("---init mylog")
		var t LoggingT
		t.output = os.Stdout
		t.lvl = cfg.GetCfgElem("logLevel").(string)
		t.loggers = make(map[string]*utils.Logger, 8)
		l = &t
		fmt.Println("---done mylog")
	})
	return l
}

var Module = fx.Options(fx.Provide(NewMylog))
