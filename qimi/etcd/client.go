package main

import (
	"context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

var gEndpoints []string

func init() {
	gEndpoints = make([]string, 0)
	gEndpoints = append(gEndpoints, "localhost:2379")
}

func main() {
	getLogagent()
}

func getLogagent() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   gEndpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("connect etcd error:", err)
	}
	log.Println("connect success")
	defer client.Close()

	key := "test1"
	// this will get a context which will unuseless afer 15*time.Second
	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	ctx := context.Background()
	resp, err := client.Get(ctx, key)
	if err != nil {
		log.Fatalf("get key:%s error:%v", key, err)
	}
	log.Println("resp is ", resp)
	// cancel()

	for _, ev := range resp.Kvs {
		log.Printf("key:%s value:%s", ev.Key, ev.Value)
	}

	// putResp, err := client.Put(ctx, "/test/lock/127.0.0.1", "1")
	// log.Printf("putresp:%v", putResp)

	log.Printf("Test4Lock start")
	Test4Lock(client, ctx, "/test/lock/127.0.0.1")
	log.Printf("Test4Lock end")

	time.Sleep(5 * time.Second)
	if err = etcdMutex.Lock(); err == nil {
		// if err = etcdMutex.LockWithTimeOut(); err == nil {
		log.Printf("get lock again")
		time.Sleep(15 * time.Second)
		etcdMutex.Unlock()
	} else {
		log.Printf("get lock error:%v", err)
	}

	// for {
	// 	watch := client.Watch(ctx, key)
	// 	for wresp := range watch {
	// 		log.Println("watch resp is ", wresp)
	// 		for _, wev := range wresp.Events {
	// 			log.Printf("type:%s, key:%q value:%q\n", wev.Type, wev.Kv.Key, wev.Kv.Value)
	// 		}
	// 	}
	// }

	log.Print("try to watch")
	watch := client.Watch(ctx, key)
	for {
		select {
		case wresp := <-watch:
			err := wresp.Err()
			if err != nil {
				log.Fatal("watch get error:", err)
			}
			for _, wev := range wresp.Events {
				log.Printf("type:%s, key:%q value:%q\n", wev.Type, wev.Kv.Key, wev.Kv.Value)
			}
		}
	}
}

type EtcdMutex struct {
	s *concurrency.Session
	m *concurrency.Mutex
}

var etcdMutex *EtcdMutex

func Test4Lock(client *clientv3.Client, ctx context.Context, key string) (mutex *EtcdMutex, err error) {
	mutex = &EtcdMutex{}
	mutex.s, err = concurrency.NewSession(client)
	if err != nil {
		log.Fatalf("concurrency.NewSession error:%v", err)
		return
	}

	mutex.m = concurrency.NewMutex(mutex.s, key)
	etcdMutex = mutex

	err = etcdMutex.LockWithTimeOut()
	if err == nil {
		for index := 0; index < 10; index++ {
			log.Print("get lock, happy")
			time.Sleep(time.Second)
		}
	} else {
		log.Printf("get lock fail:%v", err)
	}
	etcdMutex.Unlock()

	return
}

func (mutex *EtcdMutex) LockWithTimeOut() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //设置5s超时
	defer cancel()
	if err = mutex.m.Lock(ctx); err != nil {
		// err = errors.Wrap(err, "获取分布式锁失败")
	}
	return
}

func (mutex *EtcdMutex) Lock() (err error) {
	err = mutex.m.Lock(context.TODO())
	return
}

func (mutex *EtcdMutex) Unlock() (err error) {
	err = mutex.m.Unlock(context.TODO())
	return
}

func (mutex *EtcdMutex) Release() (err error) {
	err = mutex.s.Close()
	return
}