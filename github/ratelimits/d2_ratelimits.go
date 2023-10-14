package main

import (
	"github.com/juju/ratelimit"
	"log"
	"time"
)

func main2() {
	bucket := ratelimit.NewBucketWithRate(1, 5)

	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		flag := bucket.WaitMaxDuration(2, time.Second)
		if ! flag {
			log.Println(time.Now().Format("15:4:5"), "超时了")
		} else {
			log.Println(time.Now().Format("15:4:5"), "取出了2个，剩余：", bucket.Available())
		}
	}
}
