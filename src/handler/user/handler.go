package user

import (
	"github.com/jinzhu/copier"
	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"

	"github.com/phuwn/yahee/src/model"
	"github.com/phuwn/yahee/src/server"
)

func createUser(c echo.Context, u *model.User) error {
	cfg := server.GetServerCfg()
	err := cfg.Store().User.Create(c, u)
	if err != nil {
		return errors.Customize(err, 500, "create user failed")
	}
	return nil
}

// VerifyGoogleUser - verify user's google auth code, response google info of user
func VerifyGoogleUser(c echo.Context, code, redirectURL string) (*model.User, error) {
	cfg := server.GetServerCfg()
	token, err := cfg.Service().Google.GetOauth2Token(code, redirectURL)
	if err != nil {
		return nil, err
	}

	return cfg.Service().Google.GetUserGoogleInfo(token)
}

// FirstOrCreate - middleman layer to get the first record that match user's email or create new user if it doesn't exist
func FirstOrCreate(c echo.Context, u *model.User) error {
	cfg := server.GetServerCfg()
	res, err := cfg.Store().User.GetByEmail(c, u.Email)
	if err != nil {
		if errors.IsRecordNotFound(err) {
			return createUser(c, u)
		}
		return errors.Customize(err, 500, "failed to get user with email: "+u.Email)
	}

	err = copier.Copy(u, res)
	if err != nil {
		return errors.Customize(err, 500, "failed to copy value")
	}
	return nil
}

// Get - get user data from the database by the id
func Get(c echo.Context, id string) (*model.User, error) {
	cfg := server.GetServerCfg()
	res, err := cfg.Store().User.Get(c, id)
	if err != nil {
		return nil, errors.Customize(err, 404, "user not found")
	}
	return res, nil
}

// Update - update user data to the database record
func Update(c echo.Context, user *model.User) error {
	cfg := server.GetServerCfg()
	err := cfg.Store().User.Save(c, user)
	if err != nil {
		return errors.Customize(err, 500, "failed to update user")
	}
	return nil
}
