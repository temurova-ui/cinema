package user

import (
	"fmt"
	"context"
	userpb "github.com/temurova-ui/cinema/userservice/userpb/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGateway interface {
	GetUser(ctx context.Context, id int) (*userpb.GetUserResponse, error)
}

type gateway struct {
	client userpb.UserServiceClient
}

func New(address string) (UserGateway, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to userpb service: %w", err)
	}

	return &gateway{
		client: userpb.NewUserServiceClient(conn),
	}, nil
} 

func (g *gateway) GetUser(ctx context.Context, id int) (*userpb.GetUserResponse, error) {
	return g.client.GetByID(ctx, &userpb.GetUserRequest{Id: int64(id)})
}