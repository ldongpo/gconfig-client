package GRcpConfig

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
	"github.com/ldongpo/gconfig-client/kitex_gen/grpcConfig"
	"github.com/ldongpo/gconfig-client/kitex_gen/grpcConfig/grpcconfig"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	FN         = "application" //生成配置文件的名称
	Ext        = "json"        //生成配置文件的后缀
	path       = "HOME"        //生成配置文件的根目录
	pathDir    = "go-grpc-config"
	ServerHost = "GO_GRPC_CONFIG_SERVER_HOST" //服务端地址
	ENV        = "GO_GRPC_CONFIG_ENV"         //环境
	NAMESPACE  = "GO_GRPC_CONFIG_NAMESPACE"   //命名空间
	PROJECT    = "GO_GRPC_CONFIG_PROJECT"     //项目
	VERSION    = "GO_GRPC_CONFIG_VERSION"     //版本
	syncPeriod = 30
)

var (
	I    *info
	once = &sync.Once{}
)

func startInfo() {
	once.Do(initInfo)
}

func initInfo() {
	if I != nil {
		return
	}
	var err error
	I, err = createInfo()
	if err != nil {
		log.Fatalf("Fatal error GRcpConfig newInfo : %v\n", err)
	}
	err = I.sync()
	if err != nil {
		log.Fatalf("Fatal error GRcpConfig sync : %v\n", err)
	}
	go func() {
		for range time.Tick(time.Second * syncPeriod) {
			err = I.sync()
			if err != nil {
				log.Printf("sync err: %v", err)
			}
		}
	}()
}

func createInfo() (*info, error) {
	inf := &info{}
	if os.Getenv(ENV) == "" {
		return nil, errors.New("env cannot be empty")
	}
	if os.Getenv(NAMESPACE) == "" {
		return nil, errors.New("namespace cannot be empty")
	}
	if os.Getenv(PROJECT) == "" {
		return nil, errors.New("project cannot be empty")
	}
	if os.Getenv(VERSION) == "" {
		return nil, errors.New("version cannot be empty")
	}
	inf.env = os.Getenv(ENV)
	inf.namespace = os.Getenv(NAMESPACE)
	inf.project = os.Getenv(PROJECT)
	inf.version = os.Getenv(VERSION)
	if os.Getenv(ServerHost) == "" {
		return nil, errors.New("server host cannot be empty")
	}
	inf.Path = filepath.Join(os.Getenv(path), pathDir, inf.namespace, inf.project, inf.version)
	err := os.MkdirAll(inf.Path, 0755)
	if err != nil {
		return inf, err
	}
	rpcTimeout := client.WithRPCTimeout(5 * time.Second)
	connTimeout := client.WithConnectTimeout(500 * time.Millisecond)
	transport := client.WithTransportProtocol(transport.GRPC)
	host := client.WithHostPorts(os.Getenv(ServerHost))
	inf.client, err = grpcconfig.NewClient("grcpConfig", rpcTimeout, connTimeout, transport, host)
	return inf, err
}

type info struct {
	env        string
	namespace  string
	project    string
	version    string
	Path       string
	ServerHost string
	client     grpcconfig.Client
}

func (inf *info) sync() error {

	req := grpcConfig.Request{
		Env:       inf.env,
		Namespace: inf.namespace,
		Project:   inf.project,
		Version:   inf.version,
	}
	res, err := inf.client.Get(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("config response  message: %s\n", err.Error())
	}
	//log.Printf("config response: %v\n", res.Data)
	d := res.Data
	if d == "" {
		return nil
	}
	err = ioutil.WriteFile(filepath.Join(inf.Path, fmt.Sprintf("%s.%s", FN, Ext)), []byte(d), 0644)
	if err != nil {
		return err
	}
	return nil
}
