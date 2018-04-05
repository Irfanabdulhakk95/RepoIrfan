package endpoint

import (
	"context"
	"fmt"

	sv "pabriktahu/karyawan/server"
)

func (ke KaryawanEndpoint) AddKaryawanService(ctx context.Context, karyawan sv.Karyawan) error {
	_, err := ke.AddKaryawanEndpoint(ctx, karyawan)
	return err
}

func (ke KaryawanEndpoint) ReadKaryawanByNamaService(ctx context.Context, namakaryawan string) (sv.Karyawan, error) {
	req := sv.Karyawan{NamaKaryawan: namakaryawan}
	fmt.Println(req)
	resp, err := ke.ReadKaryawanByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Karyawan)
	return cus, err
}

func (ke KaryawanEndpoint) ReadKaryawanService(ctx context.Context) (sv.Karyawans, error) {
	resp, err := ke.ReadKaryawanEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Karyawans), err
}

func (ke KaryawanEndpoint) UpdateKaryawanService(ctx context.Context, kar sv.Karyawan) error {
	_, err := ke.UpdateKaryawanEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
