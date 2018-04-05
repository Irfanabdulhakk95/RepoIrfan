package main

import (
	"context"
	"fmt"
	"time"

	cli "pabriktahu/kendaraan/endpoint"
	//svc "pabriktahu/kendaraan/server"
	opt "pabriktahu/util/grpc"
	util "pabriktahu/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKendaraanClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Kendaraan
	//client.AddKendaraanService(context.Background(), svc.Kendaraan{KodeKendaraan: "003", JenisKendaraan: "DAIHATSU", NomorPlat: "123AB", CreateBy: "Admin"})

	//Get Kendaraan By NomorPlat
	//cusNama, _ := client.ReadKendaraanByNomorPlatService(context.Background(), "111Ac")
	//fmt.Println("Kendaraan based on nomorplat:", cusNama)

	//Get Kendaraan By PLAT DAERAH AB
	cuk, _ := client.ReadKendaraanByKeteranganService(context.Background(), "%AB")
	fmt.Println("all suppliers:", cuk)

	//List Kendaraan
	//cuss, _ := client.ReadKendaraanService(context.Background())
	//fmt.Println("all kendaraans:", cuss)

	//Update Kendaraan
	//client.UpdateKendaraanService(context.Background(), svc.Kendaraan{JenisKendaraan: "Monster Truck", NomorPlat: "11AC", Status: 0, UpdateBy: "Admin2", KodeKendaraan: "004"})

}
