package web

import (
	"context"
	"fmt"
	"github.com/dafsic/assistant/config"
	"github.com/dafsic/assistant/lib/mylog"
	"github.com/dafsic/assistant/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"time"
)

type Server struct {
	log *utils.Logger
	srv *http.Server
	gin *gin.Engine
}

func NewRouter(lc fx.Lifecycle, cfg config.ConfigI, log mylog.LoggingI) *Server {
	fmt.Println("---init router")
	r := gin.New()
	l := log.GetLogger("web")

	r.Use(Logger(l))
	api := cfg.GetCfgElem("api").(config.API)
	web := &Server{
		log: l,
		gin: r,
		srv: &http.Server{
			Addr: api.Address,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * time.Duration(api.Timeout),
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      r,
		},
	}

	lc.Append(fx.Hook{
		// app.start调用
		OnStart: func(ctx context.Context) error {
			// 这里不能阻塞
			go func() {
				if err := web.srv.ListenAndServe(); err != nil {
					web.log.Error(err)
				}
			}()
			return nil
		},
		// app.stop调用，收到中断信号的时候调用app.stop
		OnStop: func(ctx context.Context) error {
			web.srv.Shutdown(ctx)
			return nil
		},
	})
	fmt.Println("---done router")

	return web
}

// Module for fx
var RouterModule = fx.Options(fx.Provide(NewRouter))
