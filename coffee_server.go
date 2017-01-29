package main

import (
	"github.com/plantimals/grpctest/coffee"
	"golang.org/x/net/context"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) CoffeeService(ctx context.Context, cr *coffee.CoffeeRequest) (*coffee.CoffeeReply, error) {
	return &coffee.CoffeeReply{
		Coffee: &coffee.Coffee{
			ID:        cr.Id,
			Name:      "my coffee",
			CurrState: "heating",
			Beans: &coffee.Beans{
				ID:   1,
				Name: "death wish",
				Desc: "deadly",
			},
			History: []coffee.Transition{
				&coffee.Transition{
					ID:   1,
					From: "start",
					To:   "heating",
					Time: 1485656777,
					User: &coffee.User{
						ID:   1,
						Name: "rob",
					},
				},
			},
			CreatedAt: 1485656777,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	coffee.RegisterCoffeeServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
