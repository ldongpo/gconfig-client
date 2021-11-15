package gconfig_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/client"
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
	Ext        = "yaml"        //生成配置文件的后缀
	path       = "HOME"        //生成配置文件的根目录
	pathDir    = "go-grpc-config"
	ServerHost = "GO_GRPC_CONFIG_SERVER_HOST"
	syncPeriod = 30
)

var (
	I    *info
	once = &sync.Once{}
)

func startInfo(g *GRcpConfig) {
	once.Do(func() { initInfo(g) })
}

func initInfo(g *GRcpConfig) {
	if I != nil {
		return
	}
	var err error
	I, err = createInfo(g)
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

func createInfo(g *GRcpConfig) (*info, error) {
	inf := &info{}
	inf.env = g.Env
	inf.namespace = g.Namespace
	inf.project = g.Project
	inf.version = g.Version
	if os.Getenv(ServerHost) == "" {
		return nil, errors.New("server host cannot be empty")
	}
	inf.Path = filepath.Join(os.Getenv(path), pathDir, inf.namespace, inf.project, inf.version)
	err := os.MkdirAll(inf.Path, 0755)
	if err != nil {
		return inf, err
	}
	inf.client, err = grpcconfig.NewClient("grcpConfig", client.WithHostPorts(os.Getenv(ServerHost)))
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
