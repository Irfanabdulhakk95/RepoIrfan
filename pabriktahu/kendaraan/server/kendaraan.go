package server

import (
	"context"
)

type kendaraan struct {
	writer ReadWriter
}

func NewKendaraan(writer ReadWriter) KendaraanService {
	return &kendaraan{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *kendaraan) AddKendaraanService(ctx context.Context, kendaraan Kendaraan) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddKendaraan(kendaraan)
	if err != nil {
		return err
	}

	return nil
}

func (c *kendaraan) ReadKendaraanService(ctx context.Context) (Kendaraans, error) {
	kar, err := c.writer.ReadKendaraan()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *kendaraan) UpdateKendaraanService(ctx context.Context, kar Kendaraan) error {
	err := c.writer.UpdateKendaraan(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *kendaraan) ReadKendaraanByNomorPlatService(ctx context.Context, nomorplat string) (Kendaraan, error) {
	kar, err := c.writer.ReadKendaraanByNomorPlat(nomorplat)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
func (c *kendaraan) ReadKendaraanByNoPlatService(ctx context.Context) (Kendaraans, error) {
	kar, err := c.writer.ReadKendaraan()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
