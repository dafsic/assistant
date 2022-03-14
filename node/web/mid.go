package web

import (
	"bytes"
	"github.com/dafsic/assistant/utils"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type responseLogWriter struct {
	gin.ResponseWriter
	responseBody *bytes.Buffer
}

func (w responseLogWriter) Write(b []byte) (int, error) {
	w.responseBody.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logger access log
func Logger(l *utils.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		sid := start.UnixNano()
		ctx.Set("sessionId", sid)

		urlPath := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		method := ctx.Request.Method
		clientIP := ctx.ClientIP()
		ctx.Set("clientIP", clientIP)
		bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)

		if raw != "" {
			urlPath = urlPath + "?" + raw
		}

		//l.Infof("|%s|%s|%s|%s\n",
		//	clientIP,
		//	method,
		//	urlPath,
		//	strings.ReplaceAll(string(bodyBytes), "\n", ""),
		//)

		ctx.Set("rawData", bodyBytes)
		blw := responseLogWriter{responseBody: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = &blw //接口赋值要用地址

		ctx.Next()

		end := time.Now()
		latency := end.Sub(start)

		statusCode := ctx.Writer.Status()

		l.Infof("|%d|%v|%s|%s|%s|%s|%s\n",
			statusCode,
			latency,
			clientIP,
			method,
			urlPath,
			strings.ReplaceAll(string(bodyBytes), "\n", ""),
			blw.responseBody.Bytes(),
		)
	}
}
