package main

import (
	message "0x3f3f3f3f/pbjs_lizi/protoidl"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	engine := gin.Default()
	// 跨域处理
	engine.Use(func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	})

	engine.GET("/query", func(ctx *gin.Context) {
		var resp message.AMiniResponse
		h := md5.New()
		h.Write([]byte(time.Now().Format(time.ANSIC)))                                    // 随便写入点什么东西， 反正我也看不懂
		resp.Age = math.MaxInt64                                                          // 最大的 int64
		resp.Jwt = base64.StdEncoding.EncodeToString(h.Sum(nil))                          // 这里我是乱描述的， 是那个意思就好
		resp.Age = 18                                                                     // 心理年龄永远是年轻的
		resp.OpenId = "init_heap; init_start_up_men;load_main_fn; println('hello world')" // 初始化堆, 记载 main, 第一句永远是 hello world
		resp.Address = "on my way"                                                        // 地址是伪造的， 据说这样可以。。。
		fmt.Printf("AMiniResponse%#v\n", resp.String())
		ctx.ProtoBuf(http.StatusOK, &resp) // 使用 pb 的方式返回数据
	})

	engine.POST("/query", func(ctx *gin.Context) {
		var req message.AMiniPostRequest
		var resp message.AMiniPostResponse
		defer ctx.Request.Body.Close()
		data, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			fmt.Printf("[error] ioutil.ReadAll error:%s\n", err)
			resp.ErrCode = 500
			resp.ErrMessage = "internal error"
			ctx.ProtoBuf(http.StatusOK, &resp)
			return
		}

		if err = proto.UnmarshalMerge(data, &req); err != nil {
			fmt.Printf("[error] proto.UnmarshalMerge error:%s\n", err)
			resp.ErrCode = 400
			resp.ErrMessage = "bad params"
			ctx.ProtoBuf(http.StatusOK, &resp)
			return
		}
		resp.ErrCode = 1
		resp.ErrMessage = "success"
		for i := 0; i < rand.Intn(20); i++ {
			resp.Category = append(resp.Category, string('Z'-rand.Intn(20)))
		}
		fmt.Printf("resp:%s\n", resp.String())
		ctx.ProtoBuf(http.StatusOK, &resp)
	})


	if err := engine.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
