package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate = validator.New()
	uni      = ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
)

func init() {
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

// 定义一个结构体来接受参数
type IndexPost struct {
	QueryId int64  `json:"query_id"`
	Name    string `json:"name"`
}

// 处理 index 业务
func handleIndexPost(ctx *gin.Context) {
	var req IndexPost

	// 获取参数
	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println("error", err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "error" + err.Error()})
		return
	}

	// 校验参数
	if err := validate.Struct(&req); err != nil {
		if canTrans, ok := err.(validator.ValidationErrors); ok {
			translate := canTrans.Translate(trans)
			ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params ", "data": translate})
			return
		} else {
			fmt.Errorf("error %s\n", err)
			ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params ", "data": err.Error()})
			return
		}
	}
	// 真正处理业务
	ctx.AbortWithStatusJSON(200, req)

}

// 接受另外的一个请求参数的结构体
type OtherRequest struct {
	QueryId int64  `json:"query_id"`
	Name    string `json:"name"`
}

// 处理另外一个业务的逻辑
func handleOtherPost(ctx *gin.Context) {
	var req OtherRequest
	// 获取参数

	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println("error", err)
		ctx.AbortWithStatusJSON(400, gin.H{"error": "error" + err.Error()})
		return
	}

	// 校验参数
	if err := validate.Struct(&req); err != nil {
		if canTrans, ok := err.(validator.ValidationErrors); ok {
			translate := canTrans.Translate(trans)
			ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params ", "data": translate})
			return
		} else {
			fmt.Errorf("error %s\n", err)
			ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params ", "data": err.Error()})
			return
		}
	}
	// 真正处理业务
	ctx.AbortWithStatusJSON(200, req)
}

func init() {

}

func main() {
	e := gin.Default()
	e.POST("index", handleIndexPost)
	e.POST("/other", handleOtherPost)

	if err := e.Run(":8088"); err != nil {
		panic(err)
	}
}
