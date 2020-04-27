# 小程序使用 protobuff 

**小程序屏蔽了 google 提供 proto2js 库的某些操作, 然后世界上最大的悲剧发生了 /(ㄒoㄒ)/~~**



## 安装 proto 转 js 的工具

必要的依赖库: [ https://github.com/protobufjs/protobuf.js ]( https://github.com/protobufjs/protobuf.js ) 

主要引用的文件:  [  https://github.com/protobufjs/protobuf.js/tree/master/src ]( https://github.com/protobufjs/protobuf.js/tree/master/src )

pbjs 状态图：

![68747470733a2f2f70726f746f6275666a732e6769746875622e696f2f70726f746f6275662e6a732f746f6f6c7365742e737667](./md_image/68747470733a2f2f70726f746f6275666a732e6769746875622e696f2f70726f746f6275662e6a732f746f6f6c7365742e737667.svg)

> 命令中的 -g 选项表示全局安装

```shell
npm install -g protobufjs
```
> ps: 如果 npm 安装失败，试试国内镜像， 比如万能的淘宝 `npm config set registry https://registry.npm.taobao.org `

![npm_install_error](./md_image/npm_install_error.jpeg)


工具完成之后, 选择 `pbjs --h ` 命令大法

![pbjs__help](./md_image/pbjs__help.png)

生成的 json 描述文件 

```shell 
pbjs -t json -p . message.proto > message.json
```









