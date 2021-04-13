package repositories

import (
	"database/sql"
	"github.com/friends-management/models"
)

type IUserRepo interface {
	CreateUser(models.UserRequest) error
	GetUserList() (*models.UserListResponse, error)
}

type UserRepo struct {
	Db *sql.DB
}

func (userRepo UserRepo) CreateUser(userRequest models.UserRequest) error {
	query := `insert into users(email) values ($1)`
	_, err := userRepo.Db.Exec(query, userRequest.Email)
	return err
}

func (userRepo UserRepo) GetUserList() (*models.UserListResponse, error) {
	list := &models.UserListResponse{}
	rows, err := userRepo.Db.Query("SELECT * FROM users")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Email)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}
