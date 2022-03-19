package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/nhatlang19/go-monorepo/services/mail/mailpb"
	"google.golang.org/grpc"
)

type server struct {
	mailpb.UnimplementedMailServiceServer
}

func (s *server) SendRegisterMail(ctx context.Context, in *mailpb.RegisterMailRequest) (*mailpb.RegisterMailResponse, error) {
	log.Println("SendRegisterMail called")

	return nil, nil
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
