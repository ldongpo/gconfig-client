package GRcpConfig

import (
	"context"
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/ldongpo/gconfig-client/kitex_gen/grpcConfig"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	C *GRcpConfig
)

type GRcpConfig struct {
	v *viper.Viper
}

func init() {
	startInfo()
	var err error
	C, err = new()
	if err != nil {
		log.Fatalf("Fatal error configurator init: %v\n", err)
	}
}

// new
// @Author liangdongpo
// @Description 创建客户端
// @Date 10:34 下午 2021/11/14
// @Param
// @return
func new() (*GRcpConfig, error) {
	v := viper.New()
	v.SetConfigName(FN)
	v.SetConfigType(Ext)
	v.AddConfigPath(I.Path)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file change: %s op: %d\n", in.Name, in.Op)
	})
	return &GRcpConfig{v: v}, nil
}

// PutConfig
// @Author liangdongpo
// @Description 设置缓存
// @Date 12:33 下午 2021/11/18
// @Param
// @return
func (g *GRcpConfig) PutConfig(data string) error {
	req := grpcConfig.PutRequest{
		Env:       I.env,
		Namespace: I.namespace,
		Project:   I.project,
		Version:   I.version,
		Data:      data,
	}
	res, err := I.client.Put(context.Background(), &req)
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return errors.New(res.Msg)
	}
	return nil
}

// DelConfig
// @Author liangdongpo
// @Description 删除缓存**慎用，会把此环境、命名空间、项目、版本条件的缓存都删除
// @Date 12:35 下午 2021/11/18
// @Param
// @return
func (g *GRcpConfig) DelConfig() error {
	req := grpcConfig.DelRequest{
		Env:       I.env,
		Namespace: I.namespace,
		Project:   I.project,
		Version:   I.version,
	}
	res, err := I.client.Del(context.Background(), &req)
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return errors.New(res.Msg)
	}
	return nil
}

// GetAll
// @Author liangdongpo
// @Description 获取所有配置：string 类型
// @Date 12:45 下午 2021/11/18
// @Param
// @return
func (g *GRcpConfig) GetAll() (string, error) {
	req := grpcConfig.Request{
		Env:       I.env,
		Namespace: I.namespace,
		Project:   I.project,
		Version:   I.version,
	}
	res, err := I.client.Get(context.Background(), &req)
	if err != nil {
		return "", err
	}
	if res.Code == 0 {
		return res.Data, nil
	} else {
		return "", errors.New(res.Msg)
	}
}

// GetAllParam
// @Author liangdongpo
// @Description 自定义参数获取全部配置
// @Date 1:38 下午 2021/11/18
// @Param
// @return
func (g *GRcpConfig) GetAllParam(env string, namespace string, project string, version string) (string, error) {
	req := grpcConfig.Request{
		Env:       env,
		Namespace: namespace,
		Project:   project,
		Version:   version,
	}
	res, err := I.client.Get(context.Background(), &req)
	if err != nil {
		return "", err
	}
	if res.Code == 0 {
		return res.Data, nil
	} else {
		return "", errors.New(res.Msg)
	}
}
func (g *GRcpConfig) Get(key string) interface{} {
	return g.v.Get(key)
}

func (g *GRcpConfig) GetString(key string) string {
	return g.v.GetString(key)
}

func (g *GRcpConfig) GetBool(key string) bool {
	return g.v.GetBool(key)
}

func (g *GRcpConfig) GetInt(key string) int {
	return g.v.GetInt(key)
}

func (g *GRcpConfig) GetInt32(key string) int32 {
	return g.v.GetInt32(key)
}

func (g *GRcpConfig) GetInt64(key string) int64 {
	return g.v.GetInt64(key)
}

func (g *GRcpConfig) GetUint(key string) uint {
	return g.v.GetUint(key)
}

func (g *GRcpConfig) GetUint32(key string) uint32 {
	return g.v.GetUint32(key)
}

func (g *GRcpConfig) GetUint64(key string) uint64 {
	return g.v.GetUint64(key)
}

func (g *GRcpConfig) GetFloat64(key string) float64 {
	return g.v.GetFloat64(key)
}

func (g *GRcpConfig) GetTime(key string) time.Time {
	return g.v.GetTime(key)
}

func (g *GRcpConfig) GetDuration(key string) time.Duration {
	return g.v.GetDuration(key)
}

func (g *GRcpConfig) GetIntSlice(key string) []int {
	return g.v.GetIntSlice(key)
}

func (g *GRcpConfig) GetStringSlice(key string) []string {
	return g.v.GetStringSlice(key)
}

func (g *GRcpConfig) GetStringMap(key string) map[string]interface{} {
	return g.v.GetStringMap(key)
}

func (g *GRcpConfig) GetStringMapString(key string) map[string]string {
	return g.v.GetStringMapString(key)
}

func (g *GRcpConfig) GetStringMapStringSlice(key string) map[string][]string {
	return g.v.GetStringMapStringSlice(key)
}

func (g *GRcpConfig) GetSizeInBytes(key string) uint {
	return g.v.GetSizeInBytes(key)
}

func (g *GRcpConfig) AllSettings() map[string]interface{} {
	return g.v.AllSettings()
}
