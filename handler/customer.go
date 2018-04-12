package handler

import "github.com/gin-gonic/gin"

func (h *Handler) Register(c *gin.Context) {
	tes := make(map[string]interface{})
	tes["param1"] = "halo"
	tes["param2"] = "dunia"
	h.iseliaService.Customer.Register(tes)
}
