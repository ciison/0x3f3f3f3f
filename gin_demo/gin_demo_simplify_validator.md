## gin 项目工程之--使用设计模式简化参数校验的步骤

在以往的业务逻辑处理中， 我们可能是这样干的， 拿 `IndexRequest` 来举栗子
```go
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
```

如果再有另外的业务， 我们可能会还是会这么干的， 嗯， 这个简单， 不就是 ctrl + c & ctrl + v 么？ 这个好办

```go
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
```

嗯， 这样好像不怎么好吧， 好多代码都是重复的诶， 没有啊， 已经 "复用" 了很多的代码了啊。。。 这个。。。， 但是这个不是编译单元的代码复用

## 如果你已经厌倦了上诉的代码方式， 极力追求代码的最简洁， 那么下面就是为了你准备的
怎么能做到上面重复逻辑的复用呢？ 在 c ++ 里面是使用抽象类可以做到代码编译单元的复用， 那么使用 golang 怎么才能做到呢？ 嗯， 这就要使用 golang 的接口编程了。 

那么我们就定义一个接口， 假设就是叫做参数提取吧. 嗯， 我觉得这个名字够形象了， 起码不用猜， 对吧 😹

```go
type Extractor interface {
	Extract(ctx *gin.Context) (err error)
}
```

有接口， 要实现啊， 好吧， 我们用两个不同类型的 struct 来实现这个接口吧

```go
type IndexRequest struct {
	QueryId int64  `json:"query_id" validate:"required,gte=1"`
	Name    string `json:"name" validate:"required,ascii"`
}

func (c *IndexRequest) Extract(ctx *gin.Context) (err error) {
	return ctx.BindJSON(c)
}
```

```go
// 假设这个参数是否在请求头中的
type ProfileRequest struct {
	Token string `json:"-"`
}

func (c *ProfileRequest) Extract(ctx *gin.Context) (err error) {
	c.Token = ctx.Request.Header.Get("authorization")
	return err
}
```



将重复的逻辑包装成一个中间件

```go
func fnRequestFilter(impl Extractor) gin.HandlerFunc {
	var val = reflect.TypeOf(impl).Elem()
	return func(ctx *gin.Context) {
		filter := (reflect.New(val).Interface()).(Extractor)
		if err := filter.Extract(ctx); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "bad params"})
			return
		}

		if err := validate.Struct(filter); err == nil {
			ctx.Set("__request__", filter)
			return
		} else {
			if transErr, ok := err.(validator.ValidationErrors); ok {
				translations := transErr.Translate(trans)
				ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params", "data": translations})
				return
			} else {
				fmt.Println("error", err)
				ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params", "data": "unknown error"})
				return
			}
		}

	}
}
```



把中间件用起来：

```go
engine := gin.Default()
engine.POST("/index", fnRequestFilter(&IndexRequest{}), func(ctx *gin.Context) {
  // 在这里只写业务相关的代码
  data, _ := ctx.Get("__request__")
  ctx.JSON(200, data)
})
engine.POST("/profile", fnRequestFilter(&ProfileRequest{}), func(ctx *gin.Context) {
  // 在这里只写业务相关的代码
  data, _ := ctx.Get("__request__")
  ctx.JSON(200, data)
})
if err := engine.Run(":8089"); err != nil {
  panic(err)
}
```



贴一个工程代码：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
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
func main() {
	engine := gin.Default()
	engine.POST("/index", fnRequestFilter(&IndexRequest{}), func(ctx *gin.Context) {
		// 在这里只写业务相关的代码
		data, _ := ctx.Get("__request__")
		ctx.JSON(200, data)
	})
	engine.POST("/profile", fnRequestFilter(&ProfileRequest{}), func(ctx *gin.Context) {
		// 在这里只写业务相关的代码
		data, _ := ctx.Get("__request__")
		ctx.JSON(200, data)
	})
	if err := engine.Run("8089"); err != nil {
		panic(err)
	}
}

func fnRequestFilter(impl Extractor) gin.HandlerFunc {
	var val = reflect.TypeOf(impl).Elem()
	return func(ctx *gin.Context) {
		filter := (reflect.New(val).Interface()).(Extractor)
		if err := filter.Extract(ctx); err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "bad params"})
			return
		}

		if err := validate.Struct(filter); err == nil {
			ctx.Set("__request__", filter)
			return
		} else {
			if transErr, ok := err.(validator.ValidationErrors); ok {
				translations := transErr.Translate(trans)
				ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params", "data": translations})
				return
			} else {
				fmt.Println("error", err)
				ctx.AbortWithStatusJSON(200, gin.H{"error": "invalid params", "data": "unknown error"})
				return
			}
		}

	}
}

type Extractor interface {
	Extract(ctx *gin.Context) (err error)
}

type IndexRequest struct {
	QueryId int64  `json:"query_id" validate:"required,gte=1"`
	Name    string `json:"name" validate:"required,ascii"`
}

func (c *IndexRequest) Extract(ctx *gin.Context) (err error) {
	return ctx.BindJSON(c)
}

// 假设这个参数是否在请求头中的
type ProfileRequest struct {
	Token string `json:"-"`
}

func (c *ProfileRequest) Extract(ctx *gin.Context) (err error) {
	c.Token = ctx.Request.Header.Get("authorization")
	return err
}

```



 [代码整理工程化](https://github.com/ciison/0x3f3f3f3f/tree/master/gin_demo)

**其中请求参数实例化和参数校验的过程中使用的反射生成的对象, 要求性能的还是使用普通的方法吧** 😢

