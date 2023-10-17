package grpcclient

import (
	"context"
	"errors"

	"github.com/clicworth-com/gomicroutils/v3/genproto/auth"
)

func (a *GrpcClient) Verify(accessToken string, routeName string) (AuthInfo, error) {

	req := auth.VerifyRequest{
		Token:     accessToken,
		RouteName: routeName,
	}
	res, err := a.client.Verify(context.Background(), &req)
	if err != nil {
		return AuthInfo{}, err
	}

	if res.IsAuthorised {
		return AuthInfo{
			Authorised:  true,
			Name:        res.UserName,
			EmailId:     res.EmailId,
			PhoneNumber: res.PhoneNumber,
			Role:        res.Role,
		}, nil
	} else {
		return AuthInfo{}, errors.New(res.Message)
	}
}
