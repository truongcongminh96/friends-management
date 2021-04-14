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
	IsBlockedByUser(requestorId int, targetId int) (bool, string, error)
}

func (_subscribeRepo SubscribeRepo) CreateSubscribe(subscribe *models.Subscribe) error {
	query := `INSERT INTO subscribe(requestor, target) VALUES ($1, $2)`
	_, err := _subscribeRepo.Db.Exec(query, subscribe.Requestor, subscribe.Target)
	return err
}

func (_subscribeRepo SubscribeRepo) IsExistedSubscribe(requestorId int, targetId int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM subscribe WHERE requestor=$1 AND target=$2)`
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

func (_subscribeRepo SubscribeRepo) IsBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	var countRequestor int
	var countTarget int
	var message string

	query := `SELECT COUNT(*) FROM blocks WHERE requestor=$1 AND target=$2`

	err := _subscribeRepo.Db.QueryRow(query, requestorId, targetId).Scan(&countRequestor)
	if err != nil {
		message = "Error"
		return true, message, err
	}
	if countRequestor > 0 {
		message = "You are block the target"
		return true, message, nil
	}

	query = `SELECT COUNT(*) FROM blocks WHERE requestor=$2 AND target=$1`
	err = _subscribeRepo.Db.QueryRow(query, requestorId, targetId).Scan(&countTarget)
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
