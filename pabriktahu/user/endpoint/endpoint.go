package endpoint

import (
	"context"

	svc "git.bluebird.id/pabrikTahu/user/server"

	kit "github.com/go-kit/kit/endpoint"
)

type UserEndpoint struct {
	AddUserEndpoint       kit.Endpoint
	ReadUserByKetEndpoint kit.Endpoint
	ReadUserEndpoint      kit.Endpoint
	UpdateUserEndpoint    kit.Endpoint
}

func NewUserEndpoint(service svc.UserService) UserEndpoint {
	addUserEp := makeAddUserEndpoint(service)
	readUserByKetEp := makeReadUserByKetEndpoint(service)
	readUserEp := makeReadUserEndpoint(service)
	updateUserEp := makeUpdateUserEndpoint(service)
	return UserEndpoint{
		AddUserEndpoint:       addUserEp,
		ReadUserByKetEndpoint: readUserByKetEp,
		ReadUserEndpoint:      readUserEp,
		UpdateUserEndpoint:    updateUserEp,
	}
}

func makeAddUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		err := service.AddUserService(ctx, req)
		return nil, err
	}
}

func makeReadUserByKetEndpoint(service svc.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		result, err := service.ReadUserByKetService(ctx, req.Keterangan)
		return result, err
	}
}

func makeReadUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadUserService(ctx)
		return result, err
	}
}

func makeUpdateUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		err := service.UpdateUserService(ctx, req)
		return nil, err
	}
}
