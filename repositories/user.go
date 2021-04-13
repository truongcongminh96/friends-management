package repositories

import (
	"database/sql"
	"fmt"
	"github.com/friends-management/models"
	"strings"
)

type IUserRepo interface {
	CreateUser(user *models.User) error
	IsExistedUser(email string) (bool, error)
	GetUserIDByEmail(email string) (int, error)
	GetEmailsByIDs(listFriendId []int) ([]string, error)
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
	var isExist bool
	err := _userRepo.Db.QueryRow(query, email).Scan(&isExist)
	if err != nil {
		return true, err
	}
	if isExist {
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

func (_userRepo UserRepo) GetEmailsByIDs(listFriendId []int) ([]string, error) {
	if len(listFriendId) == 0 {
		return []string{}, nil
	}
	stringIDs := make([]string, len(listFriendId))
	for index, id := range listFriendId {
		stringIDs[index] = fmt.Sprintf("%v", id)
	}
	query := fmt.Sprintf(`SELECT email FROM users WHERE id in (%v)`, strings.Join(stringIDs, ","))
	rows, err := _userRepo.Db.Query(query)

	var emails []string

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}
	return emails, nil
}
