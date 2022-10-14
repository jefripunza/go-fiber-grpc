package communications

import (
	"context"
	"fmt"
	"go-fiber-grpc/proto"
	"os"

	"google.golang.org/grpc"
)

func BasicAdd(a uint64, b uint64) (int64, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", os.Getenv("BASIC_SERVICE")), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewBasicServiceClient(conn)

	var response int64
	var err_msg error
	req := &proto.Request{A: int64(a), B: int64(b)}
	if res, err := client.Add(context.Background(), req); err == nil {
		response = res.Result
	} else {
		err_msg = err
	}

	return response, err_msg
}

func BasicMultiply(a uint64, b uint64) (int64, error) {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", os.Getenv("BASIC_SERVICE")), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewBasicServiceClient(conn)

	var response int64
	var err_msg error
	req := &proto.Request{A: int64(a), B: int64(b)}
	if res, err := client.Multiply(context.Background(), req); err == nil {
		response = res.Result
	} else {
		err_msg = err
	}

	return response, err_msg
}
