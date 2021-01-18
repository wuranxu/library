package conf

import (
	rds "github.com/go-redis/redis"
	"github.com/wuranxu/library/log"
	"sync"
	"time"
)

var (
	once sync.Once
	Conf = new(Config)
)

type EtcdConfig struct {
	Endpoints   []string      `json:"endpoints"`
	DialTimeout time.Duration `json:"dial-timeout"`
	Scheme      string        `json:"scheme"`
}

type Config struct {
	Etcd     EtcdConfig  `json:"etcd"`
	Database SqlConfig   `json:"database"`
	Redis    rds.Options `json:"redis"`
	Scheme   string      `json:"scheme"`
}

type YamlConfig struct {
	Service string        `yaml:"service"`
	Version string        `yaml:"version"`
	Port    int           `yaml:"port"`
	Method  map[string]Md `yaml:"method"`
}

type Md struct {
	NoAuth bool `yaml:"no_auth"`
}

func Init(file string) {
	log.Info("本机环境: ", DEFAULTENV)
	var err error
	once.Do(func() {
		err = ParseConfig(file, Conf, DEFAULTENV)
		if err != nil {
			log.Fatalf("获取配置出错, error: %v", err)
		}
	})
}
