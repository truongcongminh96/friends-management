package repositories

import (
	"database/sql"
	"github.com/friends-management/models"
)

type IUserRepo interface {
	CreateUser(models.UserRequest) error
}

type UserRepo struct {
	Db *sql.DB
}

func (_self UserRepo) CreateUser(userRequest models.UserRequest) error {
	query := `insert into users(email) values ($1)`
	_, err := _self.Db.Exec(query, userRequest.Email)
	return err
}
