package core

import (
	"comment/comment_service/service"
	"comment/global"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Grpc() {
	go InitGrpc()
}

func InitGrpc() {
	global.Grpc = &global.Config.Grpc

	lis, err := net.Listen("tcp", global.Grpc.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	service.RegisterMessageServiceServer(grpcServer, service.NewCommentMessageServer())

	fmt.Println("grpc server start")
	grpcServer.Serve(lis)
}
