package conf

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/ismdeep/jwt"
	"github.com/sethvargo/go-envconfig"
)

type basicCfg struct {
	Server struct {
		Bind string `env:"SERVER_BIND,default=127.0.0.1:9000"`
		Mode string `env:"SERVER_MODE,default=debug"`
	}

	DB struct {
		Dialect string `env:"DB_DIALECT,default=sqlite"`
		DSN     string `env:"DB_DSN,default=data.db"`
	}

	System struct {
		Data string `env:"SYSTEM_DATA,default=$HOME/Documents/go-rest-api-demo-data"`
		Etcd struct {
			Endpoint string `env:"SYSTEM_ETCD_ENDPOINT,default=127.0.0.1:2379"`
		}
	}

	Log struct {
		Encoder string
		Level   string
	}
}

type systemCfg struct {
	JWT jwt.Config
}

// Basic config
var Basic basicCfg

// System config
var System systemCfg

// init basic cfg
func init() {
	if err := envconfig.Process(context.Background(), &Basic); err != nil {
		panic(err)
	}

	if Basic.System.Data == "" {
		panic("please set environment SYSTEM_DATA")
	}

	if err := os.MkdirAll(Basic.System.Data, 0750); err != nil {
		panic(err)
	}
}

// init system cfg
func init() {
	System.JWT.Key = uuid.New().String()
	System.JWT.Expire = "72h"
	initJSONMarshal("jwt", System.JWT)

	mustJSONUnmarshal("jwt", &System.JWT)
}
