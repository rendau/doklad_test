package user

import (
	"github.com/rendau/doklad_test/user/model"
)

//go:generate mockery --name UserRepoI
type UserRepoI interface {
	GetUserByID(id int) (*model.User, error)
}
