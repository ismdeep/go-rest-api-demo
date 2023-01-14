package conf

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	clientV3 "go.etcd.io/etcd/client/v3"
)

func mustGetEtcdCli() *clientV3.Client {
	cli, err := clientV3.New(clientV3.Config{
		Endpoints:   []string{Basic.System.Etcd.Endpoint},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return cli
}

func initJSONMarshal(key string, value interface{}) {
	cli := mustGetEtcdCli()
	defer func() {
		_ = cli.Close()
	}()

	// if is exists
	resp, err := cli.Get(context.Background(), key)
	if err != nil {
		panic(err)
	}
	if resp.Count > 0 {
		return
	}

	// write init data
	content, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	if _, err := cli.Put(context.Background(), key, string(content)); err != nil {
		panic(err)
	}
}

func mustJSONUnmarshal(key string, value interface{}) {
	cli := mustGetEtcdCli()
	defer func() {
		_ = cli.Close()
	}()

	resp, err := cli.Get(context.Background(), key)
	if err != nil {
		panic(err)
	}
	if resp.Count <= 0 {
		panic(fmt.Errorf("data not exists. [%v]", key))
	}

	for _, kv := range resp.Kvs {
		if string(kv.Key) == key {
			if err := json.Unmarshal(kv.Value, value); err != nil {
				panic(err)
			}
			return
		}
	}
}
