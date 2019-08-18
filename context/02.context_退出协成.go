package main

import (
	"context"
	"fmt"
	"time"
)

func Perfrom(ctx context.Context) {
	for {
		fmt.Println("计算数据")
		fmt.Println("返回结果")

		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
			fmt.Println("阻塞1s")
		}
	}
}

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, time.Hour)
	ctx, cancel :=context.WithCancel(ctx)

	go Perfrom(ctx)
	time.Sleep(time.Second * 5)
	cancel()
}
