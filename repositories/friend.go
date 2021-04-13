package repositories

import (
	"database/sql"
	"github.com/friends-management/models"
)

type FriendRepo struct {
	Db *sql.DB
}

type IFriendRepo interface {
	CreateFriend(friend *models.Friend) error
	IsExistedFriend(userID1 int, userID2 int) (bool, error)
}

func (_friendRepo FriendRepo) CreateFriend(friend *models.Friend) error {
	query := `INSERT INTO friends(user1, user2) VALUES ($1, $2)`
	_, err := _friendRepo.Db.Exec(query, friend.User1, friend.User2)
	return err
}

func (_friendRepo FriendRepo) IsExistedFriend(userID1 int, userID2 int) (bool, error) {
	query := `SELECT EXISTS(SELECT * FROM friends WHERE (user1=$1 AND user2=$2)
			 				UNION
			  				SELECT * FROM friends WHERE (user2=$1 AND user1=$2)
			 			   )`
	var isExist bool
	err := _friendRepo.Db.QueryRow(query, userID1, userID2).Scan(&isExist)
	if err != nil {
		return true, err
	}
	if isExist {
		return true, nil
	}
	return false, nil
}
