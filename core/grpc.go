package core

import (
	"comment/comment_service/service"
	"comment/global"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func GrpcServer() {
	go InitGrpcServer()
}

func InitGrpcServer() {
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

func GrpcClient() *grpc.ClientConn {
	return InitGrpcClient()
}

func InitGrpcClient() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(global.Grpc.Addr, opts...)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return conn
}
