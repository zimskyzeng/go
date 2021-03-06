package main

import (
	"github.com/juju/ratelimit"
	"log"
	"time"
)

/*
	NewBucket
 */

var (
	// 填充1个token的时间间隔
	fillInterval time.Duration = time.Millisecond * 1000
	// 容量
	capacity int64 = 10
)

func consumer(count int64, bucket *ratelimit.Bucket) {
	bucket.Wait(1)
	log.Println(time.Now().Format("2006-1-2 15:4:5"), ":取出1个")
	log.Println(time.Now().Format("2006-1-2 15:4:5"), "剩余", bucket.Available())
}

func main(){
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	ticker := time.NewTicker(time.Millisecond * 500)

	for range ticker.C {
		consumer(1, bucket)
	}
}