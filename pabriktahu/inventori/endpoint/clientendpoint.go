package endpoint

import (
	"context"
	"fmt"

	sv "pabriktahu/inventori/server"
)

func (ke InventoriEndpoint) AddInventoriService(ctx context.Context, inventori sv.Inventori) error {
	_, err := ke.AddInventoriEndpoint(ctx, inventori)
	return err
}

func (ke InventoriEndpoint) ReadInventoriByNamaInventoriService(ctx context.Context, namainventori string) (sv.Inventori, error) {
	req := sv.Inventori{NamaInventori: namainventori}
	fmt.Println(req)
	resp, err := ke.ReadInventoriByNamaInventoriEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Inventori)
	return cus, err
}

func (ke InventoriEndpoint) ReadInventoriService(ctx context.Context) (sv.Inventoris, error) {
	resp, err := ke.ReadInventoriEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Inventoris), err
}

func (ke InventoriEndpoint) UpdateInventoriService(ctx context.Context, kar sv.Inventori) error {
	_, err := ke.UpdateInventoriEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
