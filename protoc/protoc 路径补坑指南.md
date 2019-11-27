## protoc 路径补坑指南

`protoc --proto_path=IMPORT_PATH --go_out=DST_DIR 　path/file.proto`

**--proto_path** : 指定导入文件的搜索路径, 若不指定, 则为当前路径

**--go_out** : 指定生成 go 文件的路径

**path/file.proto** 定义 proto 结构的文件路径 **表示要编译的 proto 文件的路径**

#### 编译简单的 proto 文件

`protoc --go_out=plugins=grpc:. *.proto` 

这个命令表示在当前文件夹下所有的 **proto** 文件, 使用的是 *grpc* 的插件, 并把生成的文件保存在当前的文件夹

```shell
ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request (master)
$ ls
request.proto

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request (master)
$ pwd
/c/Users/cseyh/Desktop/0x3f3f3f3f/protoc/proto/request

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request (master)
$ protoc --go_out=plugins=grpc:. *.proto

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request (master)
$ ls
request.pb.go  request.proto

```

**划重点了** : 执行 `protoc` 的命令的路径是 `/Desktop/0x3f3f3f3f/protoc/proto/request` 所以 `protoc` 是在当前的路径作为根路径搜索 需要 **import | 编译 proto 文件** 

例如: 

```shell
ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request (master)
$ protoc --proto_path=. --go_out=plugins=grpc:./newpath/ ./request.proto

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request (master)
$ ls
newpath/  request.pb.go  request.proto

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request (master)
$ cd newpath/

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request/newpath (master)
$ ls
request.pb.go

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/request/newpath (master)
```

`protoc --proto_path=. --go_out=plugins=grpc:./newpath/ ./request.proto`

这个命令相当于就是把 生成的 go 文件放入到 当前路径下的 `newpath/` 文件夹下, 当前的路径是 `~/Desktop/0x3f3f3f3f/protoc/proto/request` 这就相当于新生成的 go 文件保存到 `~/Desktop/0x3f3f3f3f/protoc/proto/request/newpath` 下

#### 编译复杂的 proto 文件(带有import 其他路径的 proto 文件)



**request**: 结构体的定义

```proto
syntax="proto3";


package request;

// 定义请求结构体
message Request {

        bytes query = 1;    // 请求参数
        bytes wxOpenId = 2; // 微信平台提供的 id
        int32 code  = 3;    // 请求码
}

```

**response**: 结构体的定义

```proto
syntax="proto3";

package response;


message Response {
        bytes body = 1 ; // 请求返回体
        repeated bytes header = 2 ; // 返回的请求头
        int32 code = 3 ; // 返回的状态码
}


```

**index** : 结构体的定义

```proto
syntax= "proto3";

package index;

import "request/request.proto";
import "response/response.proto";


service Index{
        // 定义一个请求 index 的接口, 输入参数是 request.Request
        // 输出的参数是 response.Response
        rpc Index(request.Request) returns(response.Response){}

}

```

**路径的布局如下**:

> tree /f 这里我是在 windows powershell 下的命令, 这里是查看 目录结构

```shell
PS C:\Users\cseyh\Desktop\0x3f3f3f3f\protoc\proto> tree /f
卷 OS 的文件夹 PATH 列表
卷序列号为 FE0C-606D
C:.
├─index
│      index.proto
│
├─request
│  │  request.pb.go
│  │  request.proto
│  │
│  └─newpath
│          request.pb.go
│
└─response
        response.pb.go
        response.proto
```



**生成 带有 import 语句的 proto **

```shell
ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ ls
index/  request/  response/

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ pwd
/c/Users/cseyh/Desktop/0x3f3f3f3f/protoc/proto

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ protoc --proto_path=. --go_out=plugins=grpc:. ./index/index.proto

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)

```



#### 做一下错误的示范

> 根据错误的信息查看路径

```shell
ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ pwd
/c/Users/cseyh/Desktop/0x3f3f3f3f/protoc/proto

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ protoc --proto_path=./test --go_out=plugins=grpc:. ./index/index.proto
./test: warning: directory does not exist.
./index/index.proto: File does not reside within any path specified using --proto_path (or -I).  You must specify a --proto_path which encompasses this file.  Note that the proto_path must be an exact prefix of the .proto file names -- protoc is too dumb to figure out when two paths (e.g. absolute and relative) are equivalent (it's harder than you think).

```

* 报错的第一行提示 **./test** 不存在
* 第二行保存的信息说明 **./index/index.proto** 文件不存在 这个时候就会提示我们 使用 --proto_path 来指定 proto 文件的路径了



路径存在, 但是找不到需要编译 proto 文件的情况

```shell
ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ mkdir test

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ protoc --proto_path=./test --go_out=plugins=grpc:. ./index/index.proto
./index/index.proto: File does not reside within any path specified using --proto_path (or -I).  You must specify a --proto_path which encompasses this file.  Note that the proto_path must be an exact prefix of the .proto file names -- protoc is too dumb to figure out when two paths (e.g. absolute and relative) are equivalent (it's harder than you think).

```

* 可以考到报错的第一个行就提示我们 **./index/index.proto** 文件找不到

###### 导入文件的路径找不到的情况

我们把当前的路径切换到 `index.proto` 文件所在的路径

```shell
ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto (master)
$ cd index/

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/index (master)
$ pwd
/c/Users/cseyh/Desktop/0x3f3f3f3f/protoc/proto/index

ciison@DESKTOP-843CVG9 MINGW64 ~/Desktop/0x3f3f3f3f/protoc/proto/index (master)
$ protoc --proto_path=. --go_out=plugins=grpc:. ./index.proto
request/request.proto: File not found.
response/response.proto: File not found.
index.proto:5:1: Import "request/request.proto" was not found or had errors.
index.proto:6:1: Import "response/response.proto" was not found or had errors.
index.proto:12:19: "request.Request" is not defined.
index.proto:12:44: "response.Response" is not defined.
```

**重点看 protoc 报错的信息**

* `request/request.proto` 文件找不到
* `response/response.proto` 文件找不到
* `index.proto:5:1: Import "request/request.proto" was not found or had errors` index.proto 文件的第 5 行报错, 说引入的 `request/request.proto` 文件 找不到或者有错误
* `index.proto:6:1: Import "response/response.proto" was not found or had errors.` index.proto 文件的第 5 行报错, 说引入的 `response/response.proto` 文件 找不到或者有错误
* `index.proto:12:19: "request.Request" is not defined ` `request.Request` 这个结构体没有定义
* `index.proto:12:44: "response.Response" is not defined` `response.Response` 这个结构体没有定义

**综合上面的几个报错信息** 是 protoc 找不到对应的文件了

