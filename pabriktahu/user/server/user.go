package server

import (
	"context"
)

type user struct {
	writer ReadWriter
}

func NewUser(writer ReadWriter) UserService {
	return &user{writer: writer}
}

//Methode pada interface MahasiswaService di service.go
func (c *user) AddUserService(ctx context.Context, user User) error {
	//fmt.Println("mahasiswa")
	err := c.writer.AddUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (c *user) ReadUserByKetService(ctx context.Context, mob string) (Users, error) {
	us, err := c.writer.ReadUserByKet(mob)
	//fmt.Println(cus)
	if err != nil {
		return us, err
	}
	return us, nil
}

func (c *user) ReadUserService(ctx context.Context) (Users, error) {
	us, err := c.writer.ReadUser()
	//fmt.Println("mahasiswa", mhs)
	if err != nil {
		return us, err
	}
	return us, nil
}

func (c *user) UpdateUserService(ctx context.Context, us User) error {
	err := c.writer.UpdateUser(us)
	if err != nil {
		return err
	}
	return nil
}
