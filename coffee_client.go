package main

import (
	"github.com/plantimals/grpctest/coffee"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	cc := coffee.NewCoffeeServiceClient(conn)
	count := 0
	for {
		cr, err := cc.Coffee(context.Background(), &coffee.CoffeeRequest{Id: 1})
		if err != nil {
			log.Fatal(err)
			log.Print(cr)
		}
		count += 1
		if count%1000 == 0 {
			log.Print(count)
		}

	}
}
