package api

import (
	"elasticsearch_go/internal/dao/es"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct{}

func ReturnNewBook() Book {
	return Book{}
}

type response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func (b Book) CreateIndex(c *gin.Context) {
	e := es.ReturnNewEsOpreate()
	resp, err := e.CreateIndex()
	if err != nil {
		c.JSON(http.StatusOK, response{Code: -200, Msg: err, Data: ""})
		return
	}

	c.JSON(http.StatusOK, response{Code: -200, Msg: err, Data: resp})
}
