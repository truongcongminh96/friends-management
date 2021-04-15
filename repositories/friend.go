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
	IsExistedFriend(userId1 int, userId2 int) (bool, error)
	GetListFriendId(userId int) ([]int, error)
	IsBlockedByUser(requestorId int, targetId int) (bool, string, error)
	GetIdsBlockedUsers(userId int) ([]int, error)
	GetIdsSubscribers(userId int) ([]int, error)
	GetIdsBlockedByTarget(userId int) ([]int, error)
}

func (_friendRepo FriendRepo) CreateFriend(friend *models.Friend) error {
	query := `INSERT INTO friends(user1, user2) VALUES ($1, $2)`
	_, err := _friendRepo.Db.Exec(query, friend.User1, friend.User2)
	return err
}

func (_friendRepo FriendRepo) IsExistedFriend(userId1 int, userId2 int) (bool, error) {
	query := `SELECT EXISTS(SELECT * FROM friends WHERE (user1=$1 AND user2=$2)
			 				UNION
			  				SELECT * FROM friends WHERE (user2=$1 AND user1=$2)
			 			   )`
	var isExist bool
	err := _friendRepo.Db.QueryRow(query, userId1, userId2).Scan(&isExist)
	if err != nil {
		return true, err
	}
	if isExist {
		return true, nil
	}
	return false, nil
}

func (_friendRepo FriendRepo) GetListFriendId(userId int) ([]int, error) {
	query := `SELECT user1, user2 FROM friends WHERE user1=$1
			  UNION
              SELECT user1, user2 FROM friends WHERE user2=$1`

	var listFriendId = make([]int, 0)
	rows, err := _friendRepo.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id1, id2 int
		err := rows.Scan(&id1, &id2)
		if err != nil {
			return nil, err
		}
		if userId == id1 {
			listFriendId = append(listFriendId, id2)
		} else {
			listFriendId = append(listFriendId, id1)
		}
	}
	return listFriendId, nil
}

func (_friendRepo FriendRepo) IsBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	var countRequestor int
	var countTarget int
	var message string

	query := `SELECT COUNT(*) FROM blocks WHERE requestor=$1 AND target=$2`

	err := _friendRepo.Db.QueryRow(query, requestorId, targetId).Scan(&countRequestor)
	if err != nil {
		message = "Error"
		return true, message, err
	}
	if countRequestor > 0 {
		message = "You are block the target"
		return true, message, nil
	}

	query = `SELECT COUNT(*) FROM blocks WHERE requestor=$2 AND target=$1`
	err = _friendRepo.Db.QueryRow(query, requestorId, targetId).Scan(&countTarget)
	if err != nil {
		message = "Error"
		return true, message, err
	}
	if countTarget > 0 {
		message := "You are block by target"
		return true, message, nil
	}

	return false, message, nil
}

func (_friendRepo FriendRepo) GetIdsBlockedUsers(userId int) ([]int, error) {
	query := `SELECT requestor FROM blocks WHERE target=$1`

	var blockIds = make([]int, 0)
	rows, err := _friendRepo.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		blockIds = append(blockIds, id)
	}
	return blockIds, nil
}

func (_friendRepo FriendRepo) GetIdsBlockedByTarget(userId int) ([]int, error) {
	query := `SELECT target FROM blocks WHERE requestor=$1`

	var blockIds = make([]int, 0)
	rows, err := _friendRepo.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		blockIds = append(blockIds, id)
	}
	return blockIds, nil
}

func (_friendRepo FriendRepo) GetIdsSubscribers(userId int) ([]int, error) {
	query := `SELECT requestor FROM subscribe WHERE target=$1`

	var subscribers = make([]int, 0)
	rows, err := _friendRepo.Db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		subscribers = append(subscribers, id)
	}
	return subscribers, nil
}
