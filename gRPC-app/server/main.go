package main

import (
	pb "main/contract"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[gRPC-application] ")
}

func main() {
	//gRPCサーバーの設定
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterContractServiceServer(srv,
		&contractService{})
	//関数の実行
	log.Printf("start server on port%s", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
