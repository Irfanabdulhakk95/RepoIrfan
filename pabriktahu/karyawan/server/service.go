package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Karyawan.PabrikTahu.id"
	OnAdd     Status = 1
)

type Karyawan struct {
	NamaKaryawan string
	IdAgama      int32
	Alamat       string
	Jeniskelamin string
	Status       int32
	CreateBy     string
	UpdateBy     string
}
type Karyawans []Karyawan

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
	AddKaryawan(Karyawan) error
	ReadKaryawan() (Karyawans, error)
	UpdateKaryawan(Karyawan) error
	ReadKaryawanByNama(string) (Karyawan, error)
}

type KaryawanService interface {
	AddKaryawanService(context.Context, Karyawan) error
	ReadKaryawanService(context.Context) (Karyawans, error)
	UpdateKaryawanService(context.Context, Karyawan) error
	ReadKaryawanByNamaService(context.Context, string) (Karyawan, error)
}
