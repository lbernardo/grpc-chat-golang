package main

import (
	"context"
	"fmt"
	"log"
	"net"

	v1 "github.com/lbernardo/grpc-chat-golang/pkg/api/v1"
	"google.golang.org/grpc"
)

const port = "50051"

type Connection struct {
	stream v1.StarLord_ConnectStreamServer
	user   *v1.User
}

type server struct {
	Connection []*Connection
}

var users []v1.User

func (s *server) ConnectStream(u *v1.User, ss v1.StarLord_ConnectStreamServer) error {
	fmt.Printf("ConnectStream: %v\n", u.GetName())
	conn := &Connection{
		stream: ss,
		user:   u,
	}
	s.Connection = append(s.Connection, conn)
	return nil
}
func (s *server) SendMessage(ctx context.Context, msg *v1.Message) (*v1.Close, error) {
	fmt.Println("SendMessage", msg)
	for _, conn := range s.Connection {
		fmt.Println(conn.user.GetName(), "==", msg.GetUserTo().GetName())
		if conn.user.GetName() == msg.GetUserTo().GetName() {
			fmt.Println("Send message for ", msg.GetUserTo().GetName())
			err := conn.stream.Send(msg)
			fmt.Println(conn.stream)
			if err != nil {
				fmt.Printf("Error send message: %v\n", err)
			}
			break
		}
	}
	return &v1.Close{}, nil
}

func main() {
	fmt.Printf("Start server :%v\n", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	v1.RegisterStarLordServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
