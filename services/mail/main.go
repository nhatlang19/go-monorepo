package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nhatlang19/go-monorepo/services/mail/client"
	"github.com/nhatlang19/go-monorepo/services/mail/mailpb"
	"github.com/nhatlang19/go-monorepo/services/mail/service"
	"google.golang.org/grpc"
)

type server struct {
	mailpb.UnimplementedMailServiceServer
}

var (
	username = "nhatlang19@gmail.com"
	password = "Changeit@123"
	from     = "admin@example.com"
)

func (s *server) SendRegisterMail(ctx context.Context, in *mailpb.RegisterMailRequest) (*mailpb.RegisterMailResponse, error) {
	log.Println("SendRegisterMail called")

	mailTrap := client.NewMailTrap(username, password, from)
	registerService := service.NewRegisterService(mailTrap)

	status, err := registerService.Send(in.Email)
	if err != nil {
		log.Printf("[Error: %v", err)
	}

	resp := &mailpb.RegisterMailResponse{
		Status: status,
	}
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	mailpb.RegisterMailServiceServer(s, &server{})

	fmt.Println("Mail Server starting ....")
	err = s.Serve(listen)

	if err != nil {
		panic(err)
	}
}
