package node

import (
	"context"
	"fmt"
	"github.com/dafsic/assistant/config"
	"github.com/dafsic/assistant/utils"
	"github.com/dafsic/assistant/version"
	jsonrpc "github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
	"go.uber.org/fx"
	"net/http"
)

type AssistantI interface {
	Version(context.Context) (version.Version, error)
	SectorMsg(ctx context.Context, SectorId string, state string) (string, error)
	Pledge(ctx context.Context) (string, error)
}

type AssistantImpl struct {
	minerAPI  config.API
	deamonAPI config.API
	//cron time.Duration
}

func (a *AssistantImpl) Version(ctx context.Context) (version.Version, error) {
	return version.AssistantVersion, nil
}

func (a *AssistantImpl) SectorMsg(ctx context.Context, SectorId string, state string) (string, error) {
	return a.Pledge(ctx)
}

func (a *AssistantImpl) Pledge(ctx context.Context) (string, error) {
	authToken := a.minerAPI.Token
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	addr := a.minerAPI.Address

	var api lotusapi.StorageMinerStruct
	closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
	if err != nil {
		return "", fmt.Errorf("%w%s", err, utils.LineNo())
	}
	defer closer()

	id, err := api.PledgeSector(ctx)
	if err != nil {
		return "", fmt.Errorf("%w%s", err, utils.LineNo())
	}

	return id.Number.String(), nil
}

func NewAssistant(cfg config.ConfigI) AssistantI {
	a := &AssistantImpl{
		minerAPI:  cfg.GetCfgElem("minerapi").(config.API),
		deamonAPI: cfg.GetCfgElem("daemonapi").(config.API),
	}
	return a
}

var AssistantModule = fx.Options(fx.Provide(NewAssistant))
