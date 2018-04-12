package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tchalo/iselia"
	"github.com/tchalo/iselia/handler"
)

func NewWebRouter(is *iselia.IseliaService) *gin.Engine {
	r := gin.Default()
	webHandler := handler.NewHandler(is)
	r.GET("/healthz", webHandler.Healthz)

	customer := r.Group("/customer")
	{
		customer.GET("/", webHandler.Register)
	}
	return r
}
