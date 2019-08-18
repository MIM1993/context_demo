package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main01() {
	d := time.Now().Add(50 * time.Millisecond)
	fmt.Println(time.Now())
	fmt.Println(d)
	//接口 、函数
	ctx, cancel := context.WithDeadline(context.Background(), d)

	//注销函数
	defer cancel()

	select {
	   case <-time.After(1 * time.Second):
		  fmt.Println("overslept")
	   case <-ctx.Done():
	   	   fmt.Println(ctx.Err())
	}
}

func main02() {
	d := time.Now().Add(50 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(),d)
	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()
	select {
	case <- time.After(1 * time.Second):
		fmt.Println("overslept")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
	}
}


func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					os.Exit(0) // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}