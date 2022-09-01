package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"log"
	"miaosha/pb"
	"miaosha/service"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	logFile, err := os.OpenFile("./log/stock.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("set log file error: ", err)
	}
	log.SetOutput(logFile)

	port := flag.Int("port", 0, "the server port")
	serverType := flag.String("type", "grpc", "the server type")
	flag.Parse()

	stockService := service.NewStockService()
	grpcServer := grpc.NewServer()
	pb.RegisterStockServer(grpcServer, stockService)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("can not start server: ", err)
	}

	if *serverType == "grpc" {
		// grpc server
		fmt.Printf("start grpc server:%s", address)
		err = grpcServer.Serve(listener)
	} else {
		// rest server
		fmt.Printf("start rest server:%s", address)
		mux := runtime.NewServeMux()

		err = pb.RegisterStockHandlerServer(context.Background(), mux, stockService)
		if err != nil {
			log.Fatal("register rest handle error: ", err)
		}
		err = http.Serve(listener, mux)
	}
	if err != nil {
		log.Fatal("can not start server: ", err)
	}
}

