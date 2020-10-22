# readme
[李文周gin路由解析](https://www.liwenzhou.com/posts/Go/read_gin_sourcecode/#autoid-0-1-4)

[gin中文文档](https://www.kancloud.cn/shuangdeyu/gin_book/949411)
[多框架文档, 算法](http://topgoer.com)


# gin路由

## 路由原理
[原理](https://www.liwenzhou.com/posts/Go/read_gin_sourcecode/)
前缀树

### 路由拆分
根据业务拆分为多个文件

### gin 分组
```go
vv := router.Group("/v1")
	{
		vv.GET("/user/:name", handle.UserSave)
		vv.GET("/user", handle.UserSaveByQuery)
		vv.POST("/user/register", handle.UserRegister)
	}
```

## 中间件
gin框架涉及中间件相关有4个常用的方法，它们分别是c.Next()、c.Abort()、c.Set()、c.Get()。

## swagger


## viper
[viper-toml](http://www.fecmall.com/topic/1519) 大象

[viper-toml](https://blog.csdn.net/linux_player_c/article/details/82118837)

[viper git](https://github.com/spf13/viper)

## logrus
[使用logrus和rotatelog切分日志](https://blog.csdn.net/qianghaohao/article/details/104103717)

```go
package logs

import (
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/sirupsen/logrus"
    "io"
    "os"
    "sync"
    "time"
)

var logOnce sync.Once

func InitLogs(path string){

    //logFile, _ := os.OpenFile(path,os.O_CREATE|os.O_RDWR|os.O_APPEND,0644)
    logOnce.Do(func() {
        logrus.SetLevel(logrus.DebugLevel)                          // 级别
        logrus.SetFormatter(&logrus.JSONFormatter{
            //ForceColors: false,                                  // 禁止颜色
            TimestampFormat: "2006-01-02 15:05:06.000",             // 时间格式化
        })
        logrus.SetReportCaller(true)                        // 行号

        /*
           日志轮转相关函数
           `WithLinkName` 为最新的日志建立软连接
           `WithRotationTime` 设置日志分割的时间，隔多久分割一次
           `WithMaxAge` 和 `WithRotationCount` 二者只能设置一个
           `WithMaxAge` 设置文件清理前的最长保存时间
           `WithRotationCount` 设置文件清理前最多保存的个数
        */
        writer,_:=rotatelogs.New(
            path+"%Y%m%d%H%M",
            rotatelogs.WithLinkName(path),
            rotatelogs.WithMaxAge(time.Duration(180)*time.Second),
            rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
        )
        logrus.SetOutput(io.MultiWriter(os.Stdout,writer))              // 控制台和文件打印
    })
}

```

[logrus自定义格式](https://blog.csdn.net/chen09122763/article/details/105179886/)


## jwt-go
#### 生成RSA256公钥和私钥  两种方式
使用第一种
密钥可分为1024，2048，4096等位密钥，位数不同，可加解密明文长度不一。 比如1024位密钥： 可加解密明文长度 len = 1024/8 - 11 = 117字节
```shell script
openssl genrsa -out rmt.private.key 4096
openssl rsa -in rmt.private.key -pubout -outform PEM -out rmt.pub.key
## 转换为pkcs8 openssl pkcs8 -topk8 -inform PEM -in jwtRS256.key -outform pem -nocrypt -out pkcs8.pem
```

```shell script
ssh-keygen -t rsa -b 4096 -f jwtRA265.key
openssl rsa -in jwtRS265.key -pubout -outform PEM -out jwtRS265.key.pub
```
[jwt-go使用hs256](https://www.cnblogs.com/jianga/p/12487267.html)