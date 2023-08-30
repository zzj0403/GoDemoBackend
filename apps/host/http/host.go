package http

import (
	"GoDemoBackend/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
	}
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, ins)
}
