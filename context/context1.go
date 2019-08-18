package main

import (
	"context"
	"fmt"
)

var ctx = context.Background()

func SetValue() context.Context {
	ctx1 := context.WithValue(ctx, "1", "192.168.8.14")
	ctx2 := context.WithValue(ctx1, "2", "192.168.8.14")
	ctx3 := context.WithValue(ctx2, "3", "192.168.8.14")
	return ctx3
}

func GetValue(ctx context.Context) {
	v1 := ctx.Value("1")
	v2 := ctx.Value("2")
	v3 := ctx.Value("3")

	fmt.Printf("key 1 value : %v\n", v1)
	fmt.Printf("key 2 value : %v\n", v2)
	fmt.Printf("key 3 value : %v\n", v3)
}

func main() {

	ctx1 := SetValue()
	GetValue(ctx1)

	fmt.Println("finish")
}
