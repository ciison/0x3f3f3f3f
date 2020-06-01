package controller

import "github.com/gin-gonic/gin"

type profileController struct {
}

func ProfileController() *profileController {
	return &profileController{}
}
func (c *profileController) HandlerProfile(ctx *gin.Context) {
	data, _ := ctx.Get("__request__")
	// 正常的业务处理
	ctx.XML(200, data)
}

type indexController struct {
}

func IndexController() *indexController {
	return &indexController{}
}

func (c *indexController) HandlerIndex(ctx *gin.Context) {
	data, _ := ctx.Get("__request__")
	// 正常的业务处理
	ctx.AbortWithStatusJSON(200, data)
}
