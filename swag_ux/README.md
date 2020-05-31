# swag 文档工具的使用

 [来自于](https://github.com/razeencheng/demo-go/tree/master/swaggo-gin)  
**使用 swgger 可以减少手工代码**

[swaggerV2](https://swagger.io/docs/specification/2-0/basic-structure/)

使用 swagger 需要使用以下的手工包

```shell script
# swagger cli
go get -u github.com/swaggo/swag/cmd/swag
# gin-swagger 中间件
$ go get github.com/swaggo/gin-swagger
# swagger 内置文件
$ go get github.com/swaggo/gin-swagger/swaggerFiles
```
生成 doc.go 文件
```shell script
swag init 
```
引入 doc.go 文件
```go
_ "swag_ux/docs"
```


引入必要的注释
```text
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
```

* title : 文档的标题
* version: 当前的版本
* description, termsOfService, contact ... 可写可不写的声明文件
* license.name 必须要写的文件
* host: 调试 api 的主机
* BasePath: 根路径地址


添加注释代码
```go
package  main 

type Request struct {
	Id int  `json:"id" validate:"required,gte=1"`
}

type Response struct {
	Data string `json:"data"`
}

// @Summary test index
// @Description null
// @Tags index 测试
// @Accept json
// @Param  request  body Request true "查询信息"
// @Router /index [post]
// @Success 200 {object} Response
func Post (ctx *gin.Context) {
	var req Request
	var resp Response
	if err:=ctx.BindJSON(&req);err !=nil {
		resp.Data = "bad params"
		ctx.AbortWithStatusJSON(200,resp)
		return
	}

	resp .Data =  "ok " + strconv.Itoa(req.Id)
	ctx.AbortWithStatusJSON(200,resp)
	return
}
```
* Summary，description: 描述

* Tags: 用来 api 分组的

* Accept: 接受的参数类型的， 支持表单(mpfd)和JSON(json)

* Produce: 返回的数据结构

    | Mime Type                         | 声明                 |
    | --------------------------------- | -------------------- |
    | Application/json                  | json                 |
    | Application/xml                   | xml                  |
    | text/plain                        | plain                |
    | html                              | html                 |
    | Multipart/form-data               | mpfd                 |
    | Application/x-www-form-urlencoded | X-www-form-urlencode |
    | Application/vnd.api+json          | Json-api             |
    | Application/x-json-stream         | Json-stream          |
    | Application/octet-stream          | Octet-stream         |
    | Image/png                         | png                  |
    | Image/jpeg                        | jpeg                 |
    | Image/gif                         | gif                  |
    |                                   |                      |
    |                                   |                      |
    |                                   |                      |

* Param 参数

    //@Param `1.参数名` `2.参数类型` `3.参数数据类型` `4.是否是必须的` `5.参数描述` `6.其他的属性`

    1.  参数名：参数名就是解析我们参数的名字

    2.  参数类型：（参数的类型主要有三种）

        1.  `path`: 该类型的参数直接拼接在 `URL` 中， 

            比如：`// @Param  query_id path string  true "查询的id"`

        2.  `query` 该类型参数一般是组合在URL 中， 

        3.  `formData` 该类型的参数一般是 `POST`,`PUT` 方法所用

            //@Param user formData string true "用户名" default(admin)

* 参数数据类型：

    参数的数据类型一般要支持以下的几种

    * string
    * integer(int,uint, uint32, int64)
    * Number (float32)
    * Boolean (bool)

    **注意，如果你是上传文件的可以使用 `file`, 但是参数的类型一定是 `formData`**

    ```shell
    //@Param file formData file true "文件"
    ```

* 是否是必须

    * 表明参数是否是必须要的， 必须的在文档中会使用黑体标出， 表示测试时必须要填写

* 参数的描述：

    * 就是参数的说明， 比如参数有什么用的

* 其他的属性：

    * 可以设置参数的一下额外的属性， 比如枚举， 默认值， 范围等： 

        ```shell
        // 枚举的
        // @Param enumstring query string false "string enums" Enums(A, B, C)
        // @Param enumint query int false "int enums" Enums(1,2,3)
        // @Param enumnumber query false "int enums" Enums(1.2, 1.2)
        ```

        ```shell
        //值添加范围
        // @Param string query string false "string valid" minlength(5) maxlength(10)
        // @Param int query int false "int valid" mininum(1) maxinum(10)
        
        //设置默认值
        // @Param default query string false "string default" default(A)
        
        ```

        

        ```shell
         // 组合使用的
         // @Param enumstring query string false "string enums" Enums(A, B, C) default(A)
         
         
        ```

Success

​	指定成功相应的数据， 格式为： 

```shell
// @Success 1.HTTP响应码 {2.响应参数类型} 3.响应数据类型 4.其他描述
```

*   1.  HTTP 响应码， 200， 400， 500

*   2.  响应参数的类型/3. 响应的数据类型

        返回的数据类型， 可以是自定义的类型， 可以是 json

        *   自定义的响应类型

            >   一般都是 json 类型的数据

            ```shell
            // @Success 200 {object} main.File
            ```

            其中， 模式直接用 `包名.模型` 就可以了， 如果返回的是数组怎么办？

            可以这么写：

            ```shell
            // @Success 200 {array} main.File
            ```

    3.  -

*   4. 添加一些其他的描述说明信息



Failure 同 Success

Router

指定 路由与 HTTP 方法， 格式为：

>   // @Router /path/2/handler [HTTP 方法]



[源代码](./main.go)



