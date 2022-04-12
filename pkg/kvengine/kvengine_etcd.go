package kvengine

import (
	"context"
	"time"

	"github.com/Raobian/bgofs/pkg/common/log"
	"github.com/coreos/etcd/mvcc/mvccpb"

	"go.etcd.io/etcd/clientv3"
)

const (
	Endp           = "127.0.0.1:2379"
	Timeout        = 5 * time.Second
	requestTimeout = time.Second
)

type KVEngineEtcd struct {
	cli *clientv3.Client
}

func NewETCDKV() *KVEngineEtcd {
	config := clientv3.Config{
		Endpoints:   []string{Endp},
		DialTimeout: Timeout,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		log.DFATAL("etcd connect failed, err:%v", err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	status, err := cli.Status(ctx, config.Endpoints[0])
	if err != nil {
		cli.Close()
		return nil
	} else if status == nil {
		cli.Close()
		return nil
	}

	log.DINFO("etcd connect success")

	kv := &KVEngineEtcd{
		cli: cli,
	}
	return kv
}

func (kv *KVEngineEtcd) Close() error {
	return kv.cli.Close()
}

func (kv *KVEngineEtcd) Set(key string, value []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	_, err := kv.cli.Put(ctx, key, string(value))
	return err
}

func (kv *KVEngineEtcd) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	resp, err := kv.cli.Get(ctx, key)
	if err != nil {
		log.DFATAL("get from etcd failed, err:%v\n", err)
		return nil, err
	}

	if len(resp.Kvs) == 0 {
		return nil, ENotFound
	}

	log.DINFO("resp.kvs:%d", len(resp.Kvs))
	// for _, ev := range resp.Kvs {
	// 	fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	// }
	return resp.Kvs[0].Value, nil
}

func (kv *KVEngineEtcd) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	_, err := kv.cli.Delete(ctx, key)
	return err
}

func (kv *KVEngineEtcd) List(prefix string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := kv.cli.Get(ctx, prefix, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	cancel()
	if err != nil {
		log.DFATAL("List failed %v", err)
	}
	var res []string
	for _, ev := range resp.Kvs {
		res = append(res, string(ev.Key))
	}
	return res, nil
}

func (kv *KVEngineEtcd) Watch(prefix string) {
	rch := kv.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.DINFO("Watching prefix:%s", prefix)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			log.DINFO("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			switch ev.Type {
			case mvccpb.PUT:
				break
			case mvccpb.DELETE:
				break
			default:
				break
			}
		}
	}
}
