package app

import (
	"context"
	"fmt"
	"go-fiber-grpc/proto"
	"go-fiber-grpc/src/configs"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedProductsServiceServer // server gRPC sekarang
}

func GrpcServer() {

	// Run gRPC
	grpc_port := configs.GetPortGrpcServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpc_port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	svc := grpc.NewServer()
	proto.RegisterProductsServiceServer(svc, &server{})
	reflection.Register(svc)
	fmt.Printf("gRPC listen at %v ...\n", grpc_port)
	if err := svc.Serve(lis); err != nil {
		log.Fatalf("failed to start grpc serve: %v", err)
	}

}

// ===========================================================
//-> Titipan Function

func (s *server) Detail(_ context.Context, request *proto.RequestProductDetail) (*proto.ResponseProductDetail, error) {
	id := request.GetId()
	return &proto.ResponseProductDetail{Id: id}, nil
}
