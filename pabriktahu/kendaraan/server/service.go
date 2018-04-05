package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Kendaraan.PabrikTahu.id"
	OnAdd     Status = 1
)

type Kendaraan struct {
	KodeKendaraan  string
	JenisKendaraan string
	NomorPlat      string
	Status         int32
	CreateBy       string
	UpdateBy       string
}
type Kendaraans []Kendaraan

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddKendaraan(Kendaraan) error
	ReadKendaraan() (Kendaraans, error)
	UpdateKendaraan(Kendaraan) error
	ReadKendaraanByNomorPlat(string) (Kendaraan, error)
	ReadKendaraanByNoPlat(string) (Kendaraans, error)
}

type KendaraanService interface {
	AddKendaraanService(context.Context, Kendaraan) error
	ReadKendaraanService(context.Context) (Kendaraans, error)
	UpdateKendaraanService(context.Context, Kendaraan) error
	ReadKendaraanByNomorPlatService(context.Context, string) (Kendaraan, error)
	ReadKendaraanByNoPlat(context.Context) (Kendaraans, error)
}
