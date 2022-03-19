package client

import (
	"context"
	"log"

	"github.com/nhatlang19/go-monorepo/services/mail/mailpb"
	"github.com/nhatlang19/go-monorepo/services/user/model"
	"google.golang.org/grpc"
)

var (
	server_svc = "localhost:9090"
)

type MailClient interface {
	HandleRegisterMail(model.User)
}

type mailClient struct {
}

func NewMailClient() MailClient {
	return mailClient{}
}

func (m mailClient) HandleRegisterMail(user model.User) {
	conn, err := grpc.Dial(server_svc, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := mailpb.NewMailServiceClient(conn)
	status, err := c.SendRegisterMail(context.Background(), &mailpb.RegisterMailRequest{Email: user.Email})
	if err != nil {
		panic(err)
	}

	log.Printf("[Mail] Register Email called %v", status)
}
