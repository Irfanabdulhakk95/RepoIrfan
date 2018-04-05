package endpoint

import (
	"context"
	"fmt"

	sv "pabriktahu/kendaraan/server"
)

func (ke KendaraanEndpoint) AddKendaraanService(ctx context.Context, kendaraan sv.Kendaraan) error {
	_, err := ke.AddKendaraanEndpoint(ctx, kendaraan)
	return err
}

func (ke KendaraanEndpoint) ReadKendaraanByNomorPlatService(ctx context.Context, nomorplat string) (sv.Kendaraan, error) {
	req := sv.Kendaraan{NomorPlat: nomorplat}
	fmt.Println(req)
	resp, err := ke.ReadKendaraanByNomorPlatEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Kendaraan)
	return cus, err
}

func (ke KendaraanEndpoint) ReadKendaraanService(ctx context.Context) (sv.Kendaraans, error) {
	resp, err := ke.ReadKendaraanEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Kendaraans), err
}

func (ke KendaraanEndpoint) UpdateKendaraanService(ctx context.Context, kar sv.Kendaraan) error {
	_, err := ke.UpdateKendaraanEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
