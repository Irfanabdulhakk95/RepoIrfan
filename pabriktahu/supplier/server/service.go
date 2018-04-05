package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "Supplier.PabrikTahu.id"
	OnAdd     Status = 1
)

type Supplier struct {
	KodeSupplier string
	NamaSupplier string
	Status       int32
	CreateBy     string
	UpdateBy     string
	Keterangan   string
}
type Suppliers []Supplier

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
	AddSupplier(Supplier) error
	ReadSupplier() (Suppliers, error)
	UpdateSupplier(Supplier) error
	ReadSupplierByKeterangan(string) (Suppliers, error)
}

type SupplierService interface {
	AddSupplierService(context.Context, Supplier) error
	ReadSupplierService(context.Context) (Suppliers, error)
	UpdateSupplierService(context.Context, Supplier) error
	ReadSupplierByKeteranganService(context.Context, string) (Suppliers, error)
}
