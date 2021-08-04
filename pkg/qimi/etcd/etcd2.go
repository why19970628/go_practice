package main

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

var c chan int

func init() {
	c = make(chan int)
}

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://localhost:2379"},
	})
	if err != nil {
		panic(err)
	}

	watcher := clientv3.NewWatcher(client)
	channel := watcher.Watch(context.Background(), "/foobar", clientv3.WithPrefix())
	go func() {
		for {
			select {
			case change := <-channel:
				for _, ev := range change.Events {
					log.Printf("etcd change on key; %s, type = %v", string(ev.Kv.Key), ev.Type)
				}
			}
		}
	}()

	go lockFoobar(client, 1)
	go lockFoobar(client, 2)
	<-c
	<-c
}

func lockFoobar(client *clientv3.Client, id int) {
	res, err := client.Grant(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	session, err := concurrency.NewSession(client, concurrency.WithLease(res.ID))
	if err != nil {
		panic(err)
	}

	mux := concurrency.NewMutex(session, "/foobar")

	log.Printf("trying to lock by #%d\n", id)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	if err := mux.Lock(ctx); err != nil {
		log.Printf("failed to lock #%d: %v\n", id, err)
		c <- id
		return
	}

	log.Printf("post-lock #%d (lease ID = %x) bullshit\n", id, res.ID)
	time.Sleep(10 * time.Second)
	ttl, _ := client.TimeToLive(context.TODO(), res.ID)
	log.Printf("post-post-lock-#%d-sleep. lease ttl = %v", id, ttl.TTL)
	// mux.Unlock(ctx)
	// log.Printf("post-unlock #%d bullshit\n", id)

	time.Sleep(200 * time.Millisecond)
	c <- id
}
