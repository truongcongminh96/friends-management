package repositories

import (
	"database/sql"
	"github.com/friends-management/models"
)

type IUserRepo interface {
	CreateUser(user *models.User) error
	IsExistedUser(email string) (bool, error)
	GetUserIDByEmail(email string) (int, error)
}

type UserRepo struct {
	Db *sql.DB
}

func (_userRepo UserRepo) CreateUser(user *models.User) error {
	query := `INSERT INTO users(email) VALUES ($1)`
	_, err := _userRepo.Db.Exec(query, user.Email)
	return err
}

func (_userRepo UserRepo) IsExistedUser(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`
	var exist bool
	err := _userRepo.Db.QueryRow(query, email).Scan(&exist)
	if err != nil {
		return true, err
	}
	if exist {
		return true, nil
	}
	return false, nil
}

func (_userRepo UserRepo) GetUserIDByEmail(email string) (int, error) {
	query := `SELECT id FROM users WHERE email=$1`
	var id int
	err := _userRepo.Db.QueryRow(query, email).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return id, nil
}
