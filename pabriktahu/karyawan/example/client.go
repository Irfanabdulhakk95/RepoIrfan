package main

import (
	"context"
	"fmt"
	"time"

	cli "pabriktahu/karyawan/endpoint"
	opt "pabriktahu/util/grpc"
	util "pabriktahu/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKaryawanClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Karyawan
	//client.AddKaryawanService(context.Background(), svc.Karyawan{NamaKaryawan: "Irfan", IdAgama: 1, Alamat: "Rangkas", Jeniskelamin: "L", CreateBy: "Admin"})

	//Get Karyawan By Nama
	//cusNama, _ := client.ReadKaryawanByNamaService(context.Background(), "Ari Prakoso")
	//fmt.Println("karyawan based on namakaryawan:", cusNama)

	//List Karyawan
	cuss, _ := client.ReadKaryawanService(context.Background())
	fmt.Println("all karyawans:", cuss)

	//Update Karyawan
	//client.UpdateKaryawanService(context.Background(), svc.Karyawan{IdAgama: 1, Alamat: "Ciledug", Jeniskelamin: "L", Status: 0, UpdateBy: "Admin2", NamaKaryawan: "Ari Prakoso"})

}
