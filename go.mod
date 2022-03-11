module github.com/dafsic/assistant

go 1.16

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/filecoin-project/go-jsonrpc v0.1.5
	github.com/filecoin-project/lotus v1.14.4
	github.com/gin-gonic/gin v1.7.7
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/fx v1.17.0
)

replace github.com/dafsic/assistant => ./
