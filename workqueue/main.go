package main

import (
	"fmt"
	wq "k8s.io/client-go/util/workqueue"
)

func main(){
	queue := wq.NewRateLimitingQueue(wq.DefaultControllerRateLimiter())
	defer func() {
		if ! queue.ShuttingDown() {
			queue.ShutDown()
		}
	}()

	queue.Add(struct {
		name string
	}{name: "zimsky"})
	queue.Add(struct {
		age int
	}{age: 20})

	fmt.Println(queue.Len())

	item, shutdown := queue.Get()
	if ! shutdown {
		fmt.Println(item)
		queue.Done(item)
	}

	fmt.Println(queue.Len())

	item, shutdown = queue.Get()
	if ! shutdown {
		fmt.Println(item)
		queue.Done(item)
	}
	fmt.Println(queue.Len())
}