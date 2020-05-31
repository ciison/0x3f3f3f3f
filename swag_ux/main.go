package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"net/http"
	"strconv"
	_ "swag_ux/docs"
)

type Request struct {
	Id int `json:"id" validate:"required,gte=1"`
}

type Response struct {
	Data string `json:"data"`
}

// @Summary test index
// @Description null
// @Tags index 测试
// @Accept json
// @Param  request  body Request true "查询信息"
// @Param  query_id path string  true "查询的id"
// @Router /index [post]
// @Success 200 {object} Response
func Post(ctx *gin.Context) {
	var req Request
	var resp Response
	if err := ctx.BindJSON(&req); err != nil {
		resp.Data = "bad params"
		ctx.AbortWithStatusJSON(200, resp)
		return
	}

	resp.Data = "ok " + strconv.Itoa(req.Id)
	ctx.AbortWithStatusJSON(200, resp)
	return
}

// HandleHello doc
// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hello [get]
func HandleHello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}

// HandleLogin doc
// @Summary 登陆
// @Tags 登陆
// @Description 登入
// @Accept mpfd
// @Produce json
// @Param user formData string true "用户名" default(admin)
// @Param password formData string true "密码"
// @Success 200 {string} string "{"msg":"login success"}"
// @Failure 400 {string} string "{"msg": "user or password error"}"
// @Router /login [post]
func HandleLogin(c *gin.Context) {
	user := c.PostForm("user")
	pwd := c.PostForm("password")

	c.JSON(http.StatusUnauthorized, gin.H{"msg": "user or password error", "user": user, "pwd": pwd})
}

// HandleUpload doc
// @Summary 上传文件
// @Tags 文件处理
// @Description 上传文件
// @Accept mpfd
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} main.File
// @Router /upload [post]
func HandleUpload(c *gin.Context) {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	fileCon := make([]byte, 1<<20)
	n, err := file.Read(fileCon)
	if err != nil {
		if err != io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
	}

	id++
	f := &File{ID: id, Name: fileHeader.Filename, Len: int(fileHeader.Size), Content: fileCon[:n]}
	files.Files = append(files.Files, f)
	files.Len++
	c.JSON(http.StatusOK, f)
}

var files = Files{Files: []*File{}}
var id int

// File doc
type File struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Len     int    `json:"len"`
	Content []byte `json:"-"`
}

// Files doc
type Files struct {
	Files []*File `json:"files"`
	Len   int     `json:"len"`
}

// HandleList doc
// @Summary 查看文件列表
// @Tags 文件处理
// @Description 文件列表
// @Accept mpfd
// @Produce json
// @Success 200 {object} main.Files
// @Router /list [get]
func HandleList(c *gin.Context) {
	c.JSON(http.StatusOK, files)
}

// HandleGetFile doc
// @Summary 获取某个文件
// @Tags 文件处理
// @Description 获取文件
// @Accept mpfd
// @Produce octet-stream
// @Param id path integer true "文件ID"
// @Success 200 {string} string ""
// @Router /file/{id} [get]
func HandleGetFile(c *gin.Context) {
	fid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	for _, f := range files.Files {
		if f.ID == fid {
			c.Writer.WriteHeader(http.StatusOK)
			c.Header("Access-Control-Expose-Headers", "Content-Disposition")
			c.Header("Content-Disposition", "attachment; "+f.Name)
			c.Header("Content-Type", "application/octet-stream")
			c.Header("Accept-Length", fmt.Sprintf("%d", len(f.Content)))
			c.Writer.Write(f.Content)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"msg": "no avail file"})
}

// JSONParams doc
type JSONParams struct {
	// 这是一个字符串
	Str string `json:"str"`
	// 这是一个数字
	Int int `json:"int"`
	// 这是一个字符串数组
	Array []string `json:"array"`
	// 这是一个结构
	Struct struct {
		Field string `json:"field"`
	} `json:"struct"`
}

// HandleJSON doc
// @Summary 获取JSON的示例
// @Tags JSON
// @Description 获取JSON的示例
// @Accept json
// @Produce json
// @Param param body main.JSONParams true "需要上传的JSON"
// @Success 200 {object} main.JSONParams "返回"
// @Router /json [post]
func HandleJSON(c *gin.Context) {
	param := JSONParams{}
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, param)
}

// @title Swagger Example API
// @version 1.0
// @description description of your project's api
// @termsOfService null

// @contact.name null
// @contact.url
// @contact.email

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /
func main() {
	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/hello", HandleHello)
	engine.POST("/login", HandleLogin)
	engine.POST("/upload", HandleUpload)
	engine.GET("/list", HandleList)
	engine.GET("/file/:id", HandleGetFile)
	engine.POST("/json", HandleJSON)
	engine.POST("/index", Post)

	engine.Run()
}
