# Go
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/eric-jxl/Go?color=blue&label=go&logo=go)


-   Based on go version 1.21.1 darwin/amd64 
-   Golang Software Development Kit



***


### go build


|commands|Usage|
|:---|:---:|
|`-o`|指定输出的文件名，可以带上路径，例如 *go build -o a/b/c*|
|`-i`| 安装相应的包，编译+go install|
|`-a`| 更新全部已经是最新的包的，但是对标准包不适用|
|`-n`| 把需要执行的编译命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的|
|`-pn`|  指定可以并行可运行的编译数目，默认是CPU数目|
|`-race`| 开启编译的时候自动检测数据竞争的情况，目前只支持64位的机器|
|`-v`| 打印出来我们正在编译的包名|
|`-work`| 打印出来编译时候的临时文件夹名称，并且如果已经存在的话就不要删除|
|`-x`| 打印出来执行的命令，其实就是和-n的结果类似，只是这个会执行|
|`-ccflags`| *arg list* 传递参数给5c, 6c, 8c 调用|
|`-compiler`| name 指定相应的编译器，gccgo还是gc|
|`-gccgoflags`| *arg list* 传递参数给gccgo编译连接调用|
|`-gcflags`| *arg list* 传递参数给5g, 6g, 8g 调用|
|`-installsuffix`| suffix 为了和默认的安装包区别开来，采用这个前缀来重新安装那些依赖的包，-race的时候默认已经是-installsuffix race,大家可以通过-n命令来验证|
|`-ldflags`| *flag list* 传递参数给5l, 6l, 8l 调用|
|`-tags`| *tag list* 设置在编译的时候可以适配的那些tag，详细的tag限制参考里面的 Build Constraints|


### **go build** 移除路径信息
```shell
CGO_ENABLED=0 go build -v -a -ldflags '-s -w' \
-gcflags="all=-trimpath=${PWD}" \
-asmflags="all=-trimpath=${PWD}" \
-o ./main main.go
```

### **Go代码编译成动态链接库**
```shell
go build -buildmode=c-shared -o xx.so main.go
```

### **go clean** 命令用于删除执行其他命令时产生的文件或目录

### **go mod**

```markdown
download   下载依赖包
edit  修改go.mod
init 初始化项目
vendor  将依赖复制到Vendor
tidy   拉取缺失模块，移除不用的模块
graph  打印依赖图
```


### **go bug** 
`输入此命名后会直接打开默认浏览器，显示go的github页面进行bug报告，并会自动添加系统的信息。`

`//go:embed` 打包静态资源

***
:coffee: :pizza: :basketball: :lemon: :apple: :orange:

## 鸣谢

特别感谢 [JetBrains](https://www.jetbrains.com/) 为开源项目提供免费的 [Goland](https://www.jetbrains.com/go/) 授权

![Goland](https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand_icon.svg?_gl=1*a5x9r3*_ga*NzI3Njc2MDA2LjE2ODE0MDE2NDc.*_ga_9J976DJZ68*MTY4MTg2NzY1Mi4yLjEuMTY4MTg3MTU3OS4zOS4wLjA.)
