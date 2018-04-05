package server

import (
	"context"
)

type inventori struct {
	writer ReadWriter
}

func NewInventori(writer ReadWriter) InventoriService {
	return &inventori{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *inventori) AddInventoriService(ctx context.Context, inventori Inventori) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddInventori(inventori)
	if err != nil {
		return err
	}

	return nil
}

func (c *inventori) ReadInventoriService(ctx context.Context) (Inventoris, error) {
	kar, err := c.writer.ReadInventori()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *inventori) UpdateInventoriService(ctx context.Context, kar Inventori) error {
	err := c.writer.UpdateInventori(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *inventori) ReadInventoriByNamaInventoriService(ctx context.Context, namainventori string) (Inventori, error) {
	kar, err := c.writer.ReadInventoriByNamaInventori(namainventori)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
