package web

import (
	"context"
	"encoding/json"
	"github.com/dafsic/assistant/lib/mylog"
	"github.com/dafsic/assistant/node"
	"github.com/dafsic/assistant/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	assistant node.AssistantI
	log       *utils.Logger
}

func (h *Handler) SectorMsg(c *gin.Context) {
	body, _ := c.Get("rawData")
	var req struct {
		SectorId string `json:"sector_id"`
		State    string `json:"state"`
	}

	err := json.Unmarshal(body.([]byte), &req)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(200, ErrIncorrectFormat)
		return
	}

	if req.State != "p2-done" {
		h.log.Infof("sector id: %s, state: %s", req.SectorId, req.State)
		c.JSON(200, responseSuccess("ok"))
		return
	}

	id, err := h.assistant.SectorMsg(context.Background(), req.SectorId, req.State)
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(200, ErrInternalError)
		return
	}

	c.JSON(200, responseSuccess(id))
	return
}

func (h *Handler) Pledge(c *gin.Context) {
	id, err := h.assistant.Pledge(context.Background())
	if err != nil {
		h.log.Error(err.Error())
		c.JSON(200, gin.H{"data": ErrInternalError})
		return
	}

	c.JSON(200, responseSuccess(id))
	return
}

func RegisterRoutes(h *Handler, s *Server) {
	s.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	s.gin.GET("/pledge", h.Pledge)
	s.gin.POST("/msg", h.SectorMsg)
}

func NewHandler(d node.AssistantI, l mylog.LoggingI) *Handler {
	h := &Handler{
		assistant: d,
		log:       l.GetLogger("web"),
	}
	return h
}

var HandlerModule = fx.Options(fx.Provide(NewHandler))
var Register = fx.Options(fx.Invoke(RegisterRoutes))
