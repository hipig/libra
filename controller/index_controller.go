package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/lhlyu/libra/controller/dto"
	"github.com/lhlyu/libra/service/index_service"
)

type indexCtroller struct {
	controller
}

type GenerParam struct {
	Name string `json:"name"`
}

func (c indexCtroller) GenerNames(ctx iris.Context) {
	param := &dto.GenerDto{}
	if !c.getParams(ctx, param, true) {
		return
	}
	svc := index_service.NewIndexService(ctx.Request().Context())
	ctx.JSON(svc.GenerNames(param))
	return
}
