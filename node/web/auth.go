package web

//
//import (
//	"github.com/dafsic/assistant/tools/log"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strings"
//)
//
//func AuthJWT(ctx *gin.Context) {
//	token := ctx.Request.Header.Get("Authorization")
//	if token == "" {
//		token = ctx.Request.FormValue("token")
//		if token != "" {
//			token = "Bearer " + token
//		}
//	}
//
//	if token != "" {
//		if !strings.HasPrefix(token, "Bearer ") {
//			log.Warn("missing Bearer prefix in auth header")
//			ctx.JSON(http.StatusUnauthorized, ErrAuth)
//			ctx.Abort()
//			return
//		}
//		token = strings.TrimPrefix(token, "Bearer ")
//
//		allow, err := h.Verify(ctx, token)
//		if err != nil {
//			log.Warnf("JWT Verification failed (originating from %s): %s", ctx.Request.RemoteAddr, err)
//			ctx.JSON(http.StatusUnauthorized, ErrAuth)
//			ctx.Abort()
//			return
//		}
//
//		ctx = WithPerm(ctx, allow)
//	}
//
//	ctx.Next()
//}
