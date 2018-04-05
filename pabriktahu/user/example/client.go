package main

import (
	"context"
	"fmt"
	"time"

	cli "git.bluebird.id/pabrikTahu/user/endpoint"
	opt "git.bluebird.id/pabrikTahu/util/grpc"
	util "git.bluebird.id/pabrikTahu/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCUserClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Pemasukan
	//client.AddUserService(context.Background(), svc.User{Username: "damiarltsa", Password: "Dam2511", IDKaryawan: "K002", CreateBy: "Admin"})

	//Get Customer By nama kota
	userKet, _ := client.ReadUserByKetService(context.Background(), "H%")
	fmt.Println("user based on ket:", userKet)

	//List Pemasukan
	//cuss, _ := client.ReadUserService(context.Background())
	//fmt.Println("all Users:", cuss)

	//Update Pemasukan
	//client.UpdateUserService(context.Background(), svc.User{Username: "LoveYour", Password: "34351", Status: 1, UpdateBy: "Admin2", IDKaryawan: 2})

}
