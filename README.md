# webhooks

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
# sync db
./webhooks orm syncdb webhooks

# init base user
./webhooks install -username=admin -password=123456 -email=me@xiexianbin.cn

# run
bee run
bee run -downdoc=true -gendoc=true
```


