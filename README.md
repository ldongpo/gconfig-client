# gconfig-client
基于gRPC协议，集群配置中心客户端程序

### Install

`go get github.com/ldongpo/gconfig-client`



### 添加环境变量：

用于配置中心服务端微服务地址

`export GO_GRPC_SERVER_HOST="127.0.0.1:9099"`

服务端代码（暂未上传）

#### 代码示例：

```go
package main

import (
	"fmt"
	gConfig "github.com/ldongpo/gconfig-client"
)

func main() {
	var config gConfig.GRcpConfig
	config = gConfig.GRcpConfig{
		Env:       "test",  //环境
		Namespace: "semir", //命名空间
		Project:   "ts",    //项目
		Version:   "v1.0",  //版本
	}
	c, err := config.GetClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	val := c.Get("redis.host")
	fmt.Println(val)
}

```



感谢：

[viper](https://github.com/spf13/viper)

[kitex](https://github.com/cloudwego/kitex)

