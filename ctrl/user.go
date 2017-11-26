package ctrl

import "github.com/acoshift/gowebboard/api"

func NewUserController() api.UserController {
	return &userCtrl{}
}

type userCtrl struct {
}

func (c *userCtrl) SignIn(req *api.UserSignInRequest) (*api.UserSignInResponse, error) {
	return nil, nil
}

func (c *userCtrl) SignOut(req *api.UserSignOutRequest) (*api.UserSignOutResponse, error) {
	return nil, nil
}

func (c *userCtrl) ChangePassword(req *api.UserChangePasswordRequest) (*api.UserChangePasswordResponse, error) {
	return nil, nil
}

func (c *userCtrl) Get(userID int) (*api.User, error) {
	return nil, nil
}
