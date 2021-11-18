# gconfig-client
基于gRPC协议，集群配置中心客户端程序

### Install

`go get github.com/ldongpo/gconfig-client`



### 添加环境变量：

用于配置中心服务端微服务地址

```shell
export GO_GRPC_CONFIG_SERVER_HOST="127.0.0.1:9090"

```

添加环境变量：

- GO_GRPC_CONFIG_ENV：环境
- GO_GRPC_CONFIG_NAMESPACE：命名空间
- GO_GRPC_CONFIG_PROJECT：项目
- GO_GRPC_CONFIG_VERSION：版本

```sh
export GO_GRPC_CONFIG_ENV=PROD
export GO_GRPC_CONFIG_NAMESPACE="ecsemir"
export GO_GRPC_CONFIG_PROJECT="天枢"
export GO_GRPC_CONFIG_VERSION="v1.0.1"
```





服务端代码（暂未上传）

#### 代码示例：

```go
package main

import (
	"fmt"
	"github.com/ldongpo/gconfig-client"
)

func main() {
	val := GRcpConfig.C.Get("redis.db")
	fmt.Println("db:",val)
}

```



感谢：

[viper](https://github.com/spf13/viper)

[kitex](https://github.com/cloudwego/kitex)

