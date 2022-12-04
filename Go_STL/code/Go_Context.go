package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	// 10s后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
}

func doStuff(ctx context.Context) {
	for {
		// 每一秒work一下，同时判断ctx是否被取消了，如果是则退出
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}
