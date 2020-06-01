package router

import (
	"0x3f3f3f3f/gin_demo/controller"
	"0x3f3f3f3f/gin_demo/module/request"
	"0x3f3f3f3f/gin_demo/plugin/validate/zh"
	"github.com/gin-gonic/gin"
	"reflect"
)

var (
	_engine = gin.Default()
)

func init() {
	// index handler
	{
		ctr := controller.IndexController()
		_engine.POST("/index", fnRequestFilter(&request.IndexRequest{}), ctr.HandlerIndex)
	}

	// profile handler
	{
		ctr := controller.ProfileController()
		_engine.POST("/profile", fnRequestFilter(&request.ProfileRequest{}), ctr.HandlerProfile)
	}
}

func Engine() *gin.Engine {
	return _engine
}

func fnRequestFilter(impl request.Extractor) gin.HandlerFunc {
	var val = reflect.TypeOf(impl).Elem()
	return func(ctx *gin.Context) {
		filter := (reflect.New(val).Interface()).(request.Extractor)
		if err := filter.Extract(ctx); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "bad params"})
			return
		}
		trans2Zh, ok, err := zh.ValidateAndTrans2Zh(filter)

		// 参数校验通过
		if ok {
			ctx.Set("__request__", filter)
			return
		}

		if err == nil {
			// 请求的参数不符合参数的校验规范
			ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid prams", "data": trans2Zh})
			return
		} else {
			ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid prams", "data": "but not found"})
			return
		}
		//
		//if err := validate.Struct(filter); err == nil {
		//	ctx.Set("__request__", filter)
		//	return
		//} else {
		//	if transErr, ok := err.(validator.ValidationErrors); ok {
		//		translations := transErr.Translate(trans)
		//		ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params", "data": translations})
		//		return
		//	} else {
		//		fmt.Println("error", err)
		//		ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params", "data": "unknown error"})
		//		return
		//	}
		//}

	}
}
