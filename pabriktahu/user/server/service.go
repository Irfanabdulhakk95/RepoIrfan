package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "User.PabrikTahu.id"
	OnAdd     Status = 1
)

type User struct {
	Username   string
	Password   string
	IDKaryawan string
	Status     int32
	CreatedBy  string
	UpdatedBy  string
	Keterangan string
}

type Users []User

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
	AddUser(User) error
	ReadUserByKet(string) (Users, error)
	ReadUser() (Users, error)
	UpdateUser(User) error
}

type UserService interface {
	AddUserService(context.Context, User) error
	ReadUserByKetService(context.Context, string) (Users, error)
	ReadUserService(context.Context) (Users, error)
	UpdateUserService(context.Context, User) error
}
