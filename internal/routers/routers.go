package routers

import (
	"elasticsearch_go/internal/routers/api"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	book := api.ReturnNewBook()

	apiV1 := r.Group("api/v1")
	{
		apiV1.POST("/create_index", book.CreateIndex)
	}
	return r
}
