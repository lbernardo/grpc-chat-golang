// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	v1 "github.com/lbernardo/grpc-chat-golang/pkg/api/v1"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var wait *sync.WaitGroup
var client v1.StarLordClient

func init() {
	wait = &sync.WaitGroup{}
}

func connect(u *v1.User) {
	stream, err := client.ConnectStream(context.Background(), u)
	if err != nil {
		log.Fatalf("connection failed: %v\n", err)
		return
	}
	wait.Add(1)
	go func(str v1.StarLord_ConnectStreamClient) {
		defer wait.Done()
		for {
			msg, _ := str.Recv()
			if msg != nil {
				fmt.Printf("Message: %v", msg.GetContent())
			}
		}
	}(stream)
}

func main() {
	done := make(chan int)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()

	user := &v1.User{}

	fmt.Print("Your username: ")
	fmt.Scan(&user.Name)

	client = v1.NewStarLordClient(conn)

	connect(user)

	var message string
	var username string
	fmt.Print("Say user for send message (or quit): ")
	fmt.Scan(&username)
	if username == "quit" {
		os.Exit(0)
	}
	userTo := &v1.User{
		Name: username,
	}
	fmt.Print("Message: ")
	fmt.Scan(&message)

	messageSend := &v1.Message{
		Content:  message,
		UserFrom: user,
		UserTo:   userTo,
	}

	_, err = client.SendMessage(context.Background(), messageSend)
	if err != nil {
		log.Fatalf("Error sendMessage: %v", err)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done

}
