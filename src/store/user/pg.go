package user

import (
	"github.com/labstack/echo"

	"github.com/phuwn/tools/db"

	"github.com/phuwn/yahee/src/model"
)

type userPGStore struct{}

// NewStore - create new user store
func NewStore() Store {
	return &userPGStore{}
}

func (s userPGStore) Get(c echo.Context, id string) (*model.User, error) {
	tx := db.GetTxFromCtx(c)
	var res model.User
	return &res, tx.Where("id = ?", id).First(&res).Error
}

func (s userPGStore) Create(c echo.Context, user *model.User) error {
	tx := db.GetTxFromCtx(c)
	return tx.Create(user).Error
}

func (s userPGStore) GetByEmail(c echo.Context, email string) (*model.User, error) {
	tx := db.GetTxFromCtx(c)
	u := &model.User{}
	return u, tx.Where("email = ?", email).First(u).Error
}

func (s userPGStore) Save(c echo.Context, user *model.User) error {
	tx := db.GetTxFromCtx(c)
	return tx.Save(user).Error
}
