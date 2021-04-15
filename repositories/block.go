package repositories

import (
	"database/sql"

	"github.com/friends-management/models"
)

type BlockRepo struct {
	Db *sql.DB
}

type IBlockRepo interface {
	IsExistedBlock(requestorId int, targetId int) (bool, error)
	CreateBlock(block *models.Block) error
}

func (_blockRepo BlockRepo) CreateBlock(block *models.Block) error {
	query := `INSERT INTO blocks(requestor, target) VALUES ($1, $2)`
	_, err := _blockRepo.Db.Exec(query, block.Requestor, block.Target)
	return err
}

func (_blockRepo BlockRepo) IsExistedBlock(requestorId int, targetId int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM blocks WHERE requestor=$1 AND target=$2)`
	var isExist bool
	err := _blockRepo.Db.QueryRow(query, requestorId, targetId).Scan(&isExist)
	if err != nil {
		return true, err
	}
	if isExist {
		return true, nil
	}
	return false, nil
}
