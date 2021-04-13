package repositories

import (
	"database/sql"
	"github.com/friends-management/models"
)

type SubscribeRepo struct {
	Db *sql.DB
}

type ISubscribeRepo interface {
	CreateSubscribe(subscribe *models.Subscribe) error
	IsExistedSubscribe(requestorId int, targetId int) (bool, error)
}

func (_subscribeRepo SubscribeRepo) CreateSubscribe(subscribe *models.Subscribe) error {
	query := `INSERT INTO subscriptions(requestor, target) VALUES ($1, $2)`
	_, err := _subscribeRepo.Db.Exec(query, subscribe.Requestor, subscribe.Target)
	return err
}

func (_subscribeRepo SubscribeRepo) IsExistedSubscribe(requestorId int, targetId int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM subscriptions WHERE requestor=$1 AND target=$2)`
	var isExist bool
	err := _subscribeRepo.Db.QueryRow(query, requestorId, targetId).Scan(&isExist)
	if err != nil {
		return true, err
	}
	if isExist {
		return true, nil
	}
	return false, nil
}
