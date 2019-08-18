package main

import (
	pb "main/contract"
	"context"
	"log"
)

type contractService struct{}
//サーバーの処理内容
func (s *contractService) Contract(ctx context.Context, 
	req *pb.ContractRequest) (*pb.ContractResponse, error) {
	log.Println("get Contract Request")
	return &pb.ContractResponse{
		Result: "お申し込みありがとうございました。"}, nil
}
