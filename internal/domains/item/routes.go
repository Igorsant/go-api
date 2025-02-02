package item

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type router struct {
	service Service
}

func RegisterRoutes(s Service, engine *gin.Engine) {
	r := router{service: s}
	engine.GET("/item", r.getItem)
	engine.POST("/item", r.addItem)
}

func (r router) getItem(ctx *gin.Context) {
	id := ctx.Query("id")

	item, err := r.service.Get(ctx, id)
	if err != nil {
		ctx.Error(err)
		return
	}

	marshal, err := json.Marshal(item)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.String(http.StatusOK, string(marshal))
}

func (r router) addItem(ctx *gin.Context) {

	var item Item
	ctx.ShouldBind(&item)

	item, err := r.service.Add(ctx, item)
	if err != nil {
		ctx.String(http.StatusBadRequest, "error adding")
		return
	}
	marshal, err := json.Marshal(item)
	if err != nil {
		ctx.String(http.StatusBadRequest, "error marshaling")
		return
	}
	ctx.String(http.StatusAccepted, string(marshal))
}
