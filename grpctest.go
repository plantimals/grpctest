package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/plantimals/grpctest/coffee"
	"log"
)

func main() {
	coffee := &coffee.CoffeeRequest{
		Id: *proto.Int32(1),
	}
	data, err := proto.Marshal(coffee)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(coffee)
	fmt.Println(data)
	fmt.Println("hello")
}
