package repositories

import "database/sql"

type FriendRepo struct {
	Db *sql.DB
}

type IFriendRepo interface {
	IsExistedFriend()
}