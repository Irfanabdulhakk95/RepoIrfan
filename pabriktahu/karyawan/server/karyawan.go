package server

import (
	"context"
)

type karyawan struct {
	writer ReadWriter
}

func NewKaryawan(writer ReadWriter) KaryawanService {
	return &karyawan{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *karyawan) AddKaryawanService(ctx context.Context, karyawan Karyawan) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddKaryawan(karyawan)
	if err != nil {
		return err
	}

	return nil
}

func (c *karyawan) ReadKaryawanService(ctx context.Context) (Karyawans, error) {
	kar, err := c.writer.ReadKaryawan()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}

func (c *karyawan) UpdateKaryawanService(ctx context.Context, kar Karyawan) error {
	err := c.writer.UpdateKaryawan(kar)
	if err != nil {
		return err
	}
	return nil
}

func (c *karyawan) ReadKaryawanByNamaService(ctx context.Context, namakaryawan string) (Karyawan, error) {
	kar, err := c.writer.ReadKaryawanByNama(namakaryawan)
	//fmt.Println("mahasiswa:", mhs)
	if err != nil {
		return kar, err
	}
	return kar, nil
}
