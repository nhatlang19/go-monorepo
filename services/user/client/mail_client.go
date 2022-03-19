package client

import (
	"context"

	"github.com/nhatlang19/go-monorepo/services/mail/mailpb"
	"google.golang.org/grpc"
)

var (
	server_svc = "localhost:9090"
)

func handleRegisterMail() {
	conn, err := grpc.Dial(server_svc, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := mailpb.NewMailServiceClient(conn)
	sum, err := c.SendRegisterMail(context.Background(), &mailpb.RegisterMailRequest{jwt: "jwt"})
	if err != nil {
		panic(err)
	}
}
