package handler

import (
	"github.com/gin-gonic/gin"
	is "github.com/tchalo/iselia"
	"net/http"
)

type Handler struct {
	iseliaService *is.IseliaService
}

type Meta struct {
	HTTPStatus int   `json:"http_status"`
	Total      int64 `json:"total,omitempty"`
	Limit      int   `json:"limit,omitempty"`
	Offset     int   `json:"offset,omitempty"`
}

func NewHandler(iseliaService *is.IseliaService) *Handler {
	return &Handler{
		iseliaService: iseliaService,
	}
}

func (h *Handler) Healthz(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", []byte("ok"))
}
