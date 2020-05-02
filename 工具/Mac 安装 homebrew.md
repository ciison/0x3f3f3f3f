# Mac 安装 homebrew 

[brew官网](https://brew.sh)

官网给出的安装指令 

```shell 
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

这个指定不是 `raw.githubusercontent.com` 403 错误就是 一些很奇怪的问题

**raw.githubusercontent.com** 报错的问题可以通过修改 `/etc/hosts` 文件来解决

1.  查询 [raw.githubuserconent.com](https://githubusercontent.com.ipaddress.com/raw.githubusercontent.com)的解析主机的地址

2.  将 返回的 ip 地址追加到 `/etc/hosts` 文件

    >   举个例子

    ```shell
    199.232.68.133 raw.githubusercontent.com
    ```

>   然而即使是这样还有一大堆的问题等着你去解决， 比如换回国内镜像稳点



## 国内源地址解决

*   先把官网的安装脚本下载下来

    ```shell 
    cd ~ # 切换到当前用户的根目录
    curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install >> brew_install # 把官网的安装脚本下载到本地
    ```

*   修改脚本中的镜像源地址

    ```shell 
    #!/usr/bin/ruby
    
    HOMEBREW_PREFIX = "/usr/local".freeze
    HOMEBREW_REPOSITORY = "/usr/local/Homebrew".freeze
    HOMEBREW_CACHE = "#{ENV["HOME"]}/Library/Caches/Homebrew".freeze
    HOMEBREW_OLD_CACHE = "/Library/Caches/Homebrew".freeze
    #BREW_REPO = "https://github.com/Homebrew/brew".freeze
    BREW_REPO = "https://mirrors.ustc.edu.cn/brew.git".freeze
    #CORE_TAP_REPO = "https://github.com/Homebrew/homebrew-core".freeze
    CORE_TAP_REPO = "https://mirrors.ustc.edu.cn/homebrew-core.git".freeze
    
    STDERR.print <<EOS
    Warning: The Ruby Homebrew installer is now deprecated and has been rewritten in
    Bash. Please migrate to the following command:
      /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
    
    EOS
    
    Kernel.exec "/bin/bash", "-c", '/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"'
    
    ```

    >   修改完成保存， 

*   执行安装脚本的命令

    ```shell
    /usr/bin/ruby brew_install # 执行当前文件夹下的 brew_install 文件
    ```

*   运行这段命令基本上就可以看到成功的信息了



以上内容大多来自 [csdn](https://blog.csdn.net/zyrl2012/article/details/81388326)

