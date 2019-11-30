## 从零开始写一个 HTTP 框架(二)

### 请求的路由怎么动态分发呢?

使用 正则表达式可以匹配有一定特征的字符. 但是需要怎么保存路由呢? http 包里面使用了

```go
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	es    []muxEntry // slice of entries sorted from longest to shortest.
	hosts bool       // whether any patterns contain hostnames
}

type muxEntry struct {
	h       Handler
	pattern string
}

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

> 上面这一段的代码来自 go 的 http 标准库

从上面的代码可以看出, 标准库的代码的 pattern 的路径是保存在 map 里面的, 对应的处理方法是 Handler(没有使用反射), Handler 其实就是一个接口而已. 在 go 里面, 一个类型实现了对应的接口提供的方法, 就默认是实现了该接口.


#### 在 go 的世界里, 写一个简单的 web 服务再简单不过了

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/demo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(`Content-Type`, `application/json`)
		fmt.Fprint(w, fmt.Sprintf(`{"hello":"%s"}`, time.Now().String()))
	})

	http.ListenAndServe(":8080", nil)
}

```
上面的代码就是实现了一个简单的 web 服务, 如果用户的访问路径是 `/demo` 不管用户的访问方式是什么, 返回给用户的都是 json 格式的字符串

**主要的流程是: 如果服务端发现一个请求过来, 到 `muxEntry` 查找是否有对应的路径, 如果没有发现, 返回客户端 404 错误, 如果找到对应的 `handler`, 调用对应的 `handler` 处理业务逻辑**  ==> 这个就是 go 标准库内置的 http server 的处理逻辑了

如果业务足够简单, 这样的处理是可以的, 但是在这个时代, 估计没有比 [example](http://www.example.com) 更加简单的网站了, 所以对于像 `/demo/1234`, `/demo/2345` 这种格式相似的路由, 需要使用更加高级的匹配模式来处理了


> 下一篇 说的是 `树`




