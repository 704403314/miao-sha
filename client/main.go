package main

import (
	"context"
	"flag"
	"log"
	"miaosha/pb"
	"time"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := flag.String("address", "", "the server port")
	flag.Parse()
	log.Printf("dial server  %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("can not dial server: ", err)
	}

	stockClient := pb.NewStockClient(conn)

	req := &pb.SearchStock{
		GoodsId: "123",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := stockClient.GetStock(ctx, req)
	if err != nil {
		// _, _ = status.FromError(err)
		log.Fatal("cannot get stock number")
		return
	}
	log.Printf("get stock info: %v", res)
}
