package config

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/fx"
	"reflect"
	"strings"
	"time"
)

type ConfigI interface {
	GetCfg(e string) interface{}
}

// FullNode is a full node config
type AssistantNode struct {
	LogLevel  string
	API       API
	MinerAPI  API
	DaemonAPI API
}

// API contains configs for API endpoint
type API struct {
	Address string
	Token   string
	Timeout time.Duration
}

func (a *AssistantNode) GetCfgElem(e string) interface{} {
	var cnf interface{}
	rt := reflect.TypeOf(*a)
	rv := reflect.ValueOf(*a)

	fieldNum := rt.NumField()
	for i := 0; i < fieldNum; i++ {
		if strings.ToUpper(rt.Field(i).Name) == strings.ToUpper(e) {
			cnf = rv.FieldByName(rt.Field(i).Name).Interface()
			break
		}
	}
	return cnf
}

func NewAssCfg(ctx *cli.Context) (*AssistantNode, error) {
	fmt.Println("---init assistant config")
	cp := ctx.String("config")
	c, err := FromFile(cp, DefaultAssistantNode())
	if err != nil {
		return nil, err
	}
	cfg, ok := c.(*AssistantNode)
	if !ok {
		return nil, fmt.Errorf("invalid config for assistant, got: %T", c)
	}
	return cfg, nil
}

var AssitantModule = fx.Options(fx.Provide(NewAssCfg))
