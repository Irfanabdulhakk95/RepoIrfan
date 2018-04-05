package main

import (
	"context"
	"time"

	cli "pabriktahu/inventori/endpoint"
	svc "pabriktahu/inventori/server"
	opt "pabriktahu/util/grpc"
	util "pabriktahu/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCInventoriClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Inventori
	//client.AddInventoriService(context.Background(), svc.Inventori{KodeInventori: "003", NamaInventori: "kais", Jumlah: 20, CreateBy: "Admin"})

	//Get Inventori By NamaInventori
	//cusNama, _ := client.ReadInventoriByNamaInventoriService(context.Background(), "topi")
	//fmt.Println("Inventori based on namainventori:", cusNama)

	//List Inventori
	//cuss, _ := client.ReadInventoriService(context.Background())
	//fmt.Println("all inventoris:", cuss)

	//Update Inventori
	client.UpdateInventoriService(context.Background(), svc.Inventori{NamaInventori: "krucut", Jumlah: 30, Status: 1, UpdateBy: "Admin2", KodeInventori: "001"})

	//Get Inventori By Jumlah
	//cusNama, _ := client.ReadInventoriByNamaInventoriService(context.Background(), "topi")
	//fmt.Println("Inventori based on namainventori:", cusNama)

}
