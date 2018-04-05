package main

import (
	"context"
	"fmt"
	"time"

	cli "pabriktahu/supplier/endpoint"
	// svc "pabriktahu/supplier/server"
	opt "pabriktahu/util/grpc"
	util "pabriktahu/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCSupplierClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Supplier
	// client.AddSupplierService(context.Background(), svc.Supplier{KodeSupplier: "004", NamaSupplier: "adam", Keterangan: "jakarta", CreateBy: "Admin"})

	//Get Kendaraan By Keterangan
	//cusNama, _ := client.ReadSupplierByKeteranganService(context.Background(), "S")
	//fmt.Println("Kendaraan based on nomorplat:", cusNama)

	// List Suplier Berdasarkan Keterangan
	cuk, _ := client.ReadSupplierByKeteranganService(context.Background(), "j%")
	fmt.Println("all suppliers:", cuk)

	//List Supplier
	// cuss, _ := client.ReadSupplierService(context.Background())
	// fmt.Println("all suppliers:", cuss)

	//Update Supplier
	//client.UpdateSupplierService(context.Background(), svc.Supplier{NamaSupplier: "Olgi", Status: 1, Keterangan: "jakarta", UpdateBy: "Admin1", KodeSupplier: "002"})

}
