package endpoint

import (
	"context"
	"fmt"

	sv "git.bluebird.id/pabrikTahu/user/server"
)

func (ue UserEndpoint) AddUserService(ctx context.Context, user sv.User) error {
	_, err := ue.AddUserEndpoint(ctx, user)
	return err
}

func (ke UserEndpoint) ReadUserByKetService(ctx context.Context, keterangan string) (sv.Users, error) {
	req := sv.User{Keterangan: keterangan}
	fmt.Println(req)
	resp, err := ke.ReadUserByKetEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Users)
	return cus, err
}

func (ue UserEndpoint) ReadUserService(ctx context.Context) (sv.Users, error) {
	resp, err := ue.ReadUserEndpoint(ctx, nil)
	//fmt.Println("peme resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Users), err
}

func (ue UserEndpoint) UpdateUserService(ctx context.Context, user sv.User) error {
	_, err := ue.UpdateUserEndpoint(ctx, user)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
