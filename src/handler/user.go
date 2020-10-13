package handler

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"
	"github.com/phuwn/yahee/src/handler/user"
	"github.com/phuwn/yahee/src/model"
)

func userRoutes(r *echo.Echo) {
	r.POST("/auth", signIn)
	g := r.Group("/user")
	{
		g.GET("/me", getMyInfo)
		g.PUT("/me", updateMyInfo)
	}
}

func getMyInfo(c echo.Context) error {
	id := model.GetUserIDFromCtx(c)
	u, err := user.Get(c, id)
	if err != nil {
		return err
	}
	return JSON(c, 200, u)
}

// SignInRequest - data form to sign in to auth
type SignInRequest struct {
	Code        string `json:"code"`
	RedirectURL string `json:"redirect_url"`
}

func signIn(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return errors.Customize(err, 400, "unable to read the request body")
	}

	req := &SignInRequest{}
	err = json.Unmarshal(b, req)
	if err != nil {
		return errors.Customize(err, 400, "wrong sign in data form")
	}

	u, err := user.VerifyGoogleUser(c, req.Code, req.RedirectURL)
	if err != nil {
		return err
	}

	err = user.FirstOrCreate(c, u)
	if err != nil {
		return err
	}

	jwt, err := model.GenerateJWTToken(&model.TokenInfo{UserID: u.ID}, time.Now().Add(24*time.Hour).Unix())
	if err != nil {
		return err
	}

	u.AccessToken = &jwt
	return JSON(c, 200, u)
}

// UpdateInfoRequest - data form to update user's info
type UpdateInfoRequest struct {
	Name   *string `json:"name"`
	Avatar *string `json:"avatar"`
}

func updateMyInfo(c echo.Context) error {
	id := model.GetUserIDFromCtx(c)
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return errors.Customize(err, 400, "unable to read the request body")
	}

	req := &UpdateInfoRequest{}
	err = json.Unmarshal(b, req)
	if err != nil {
		return errors.Customize(err, 400, "wrong update info form")
	}
	if req.Avatar == nil && (req.Name == nil || *req.Name == "") {
		return errors.New("no change was committed", 400)
	}
	u, err := user.Get(c, id)
	if err != nil {
		return err
	}

	if req.Avatar != nil {
		u.Avatar = *req.Avatar
	}
	if !(req.Name == nil || *req.Name == "") {
		u.Name = *req.Name
	}
	err = user.Update(c, u)
	if err != nil {
		return err
	}
	return JSON(c, 200, u)
}
