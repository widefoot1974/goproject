package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "myValue")
	ctx = context.WithValue(ctx, "myKey2", "myValue2")
	ctx = context.WithValue(ctx, 123, 123)
	ctx = context.WithValue(ctx, 123, 1234)
	ctx = context.WithValue(ctx, 123, []int{1, 2, 3, 4})
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	fmt.Println("ctx value:", ctx.Value("myKey"))
	fmt.Println("ctx value:", ctx.Value("myKey2"))
	fmt.Println("ctx value:", ctx.Value(123))
}
