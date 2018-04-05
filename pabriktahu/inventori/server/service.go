package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Inventori.PabrikTahu.id"
	OnAdd     Status = 1
)

type Inventori struct {
	KodeInventori string
	NamaInventori string
	Jumlah        int32
	Status        int32
	CreateBy      string
	UpdateBy      string
}
type Inventoris []Inventori

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
	AddInventori(Inventori) error
	ReadInventori() (Inventoris, error)
	UpdateInventori(Inventori) error
	ReadInventoriByNamaInventori(string) (Inventori, error)
}

type InventoriService interface {
	AddInventoriService(context.Context, Inventori) error
	ReadInventoriService(context.Context) (Inventoris, error)
	UpdateInventoriService(context.Context, Inventori) error
	ReadInventoriByNamaInventoriService(context.Context, string) (Inventori, error)
}
