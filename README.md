# webhooks

基于 `begoo` 实现 `webhook` 功能，使用场景如 `Github`/`Gitlab` 某个分支发生 `push` 的时候，自动触发一段自定义脚本。

## 源码安装使用


### 拉取源码

```
go get github.com/xiexianbin/webhooks.git
```


### 编译源码

```
# 更新依赖
go get -d ./...

# 编译项目
go build -v -tags "pam" -ldflags "-w"
```


### 配置

待补充。


### 运行

```
# run
bee run
bee run -downdoc=true -gendoc=true
```

### 打包

```
bee pack
bee pack -be GOOS=linux -be GOARCH=amd64
```

### 使用
