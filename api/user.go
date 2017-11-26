package api

import "fmt"

// UserController is the user controller interface
type UserController interface {
	SignIn(*UserSignInRequest) (*UserSignInResponse, error)
	SignOut(*UserSignOutRequest) (*UserSignOutResponse, error)
	ChangePassword(*UserChangePasswordRequest) (*UserChangePasswordResponse, error)
	Get(int) (*User, error)
}

// UserSignInRequest is the sign in request for user controller
type UserSignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *UserSignInRequest) Validate() error {
	if len(req.Username) == 0 {
		return fmt.Errorf("user required")
	}
	if len(req.Password) == 0 {
		return fmt.Errorf("password required")
	}
	return nil
}

// UserSignInResponse is the sign in response for user controller
type UserSignInResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

// UserSignOutRequest is the sign out request for user controller
type UserSignOutRequest struct {
	Token string `json:"token"`
}

func (req *UserSignOutRequest) Validate() error {
	if len(req.Token) == 0 {
		return fmt.Errorf("token required")
	}
	return nil
}

// UserSignOutResponse is the sign out response for user controller
type UserSignOutResponse struct {
}

// UserChangePasswordRequest type
type UserChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// UserChangePasswordResponse type
type UserChangePasswordResponse struct {
}

// User type
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
