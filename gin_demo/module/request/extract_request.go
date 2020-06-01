package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Extractor interface {
	Extract(ctx *gin.Context) (err error)
}

type IndexRequest struct {
	QueryId int64  `json:"query_id" validate:"required,gte=1"` // 校验不能为空， 要大于 1
	Name    string `json:"name" validate:"required,ascii"`     // 不能为空， 满足 ascii 编码字符
}

func (c *IndexRequest) Extract(ctx *gin.Context) (err error) {
	return ctx.BindJSON(c)
}

// 假设这个参数是否在请求头中的
type ProfileRequest struct {
	Token string `json:"-" validate:"required,min=32"` // 不能为空, 最短长度为 32
}

func (c *ProfileRequest) Extract(ctx *gin.Context) (err error) {
	c.Token = ctx.Request.Header.Get("Authorization")
	fmt.Println(c)
	return err
}
