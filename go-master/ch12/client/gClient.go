package main

import (
	"context"
	"fmt"
	"go-master/ch12/protoapi"
	"math/rand"
	"os"

	"google.golang.org/grpc"
)

var port = ":8080"

func AskingDateTime(ctx context.Context, m protoapi.RandomClient) (*protoapi.DateTime, error) {
	request := &protoapi.RequestDateTime{
		Value: "Please send me the date and time",
	}

	return m.GetDate(ctx, request)
}

func AskPass(ctx context.Context, m protoapi.RandomClient, seed int64, length int64) (*protoapi.RandomPass, error) {
	request := &protoapi.RequestPass{
		Seed:   seed,
		Length: length,
	}

	return m.GetRandomPass(ctx, request)
}

func AskRandom(ctx context.Context, m protoapi.RandomClient, seed int64, place int64) (*protoapi.RandomInt, error) {
	request := &protoapi.RandomParams{
		Seed:  seed,
		Place: place,
	}

	return m.GetRandom(ctx, request)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	// seed := int64(rand.Intn(100))

	client := protoapi.NewRandomClient(conn)

	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Date and Time:", r.Value)

	length := int64(rand.Intn(20))
	p, err := AskPass(context.Background(), client, 100, length+1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Password:", p.Password)
}
