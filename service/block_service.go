package service

import (
	"github.com/friends-management/models"
	"github.com/friends-management/repositories"
)

type BlockService struct {
	IBlockRepo repositories.IBlockRepo
}

type IBlockService interface {
	CreateBlock(block *models.Block) error
	CheckExistedBlock(requestorId int, targetId int) (bool, error)
}

func (_blockService BlockService) CheckExistedBlock(requestorId int, targetId int) (bool, error) {
	isExist, err := _blockService.IBlockRepo.IsExistedBlock(requestorId, targetId)
	return isExist, err
}

func (_blockService BlockService) CreateBlock(block *models.Block) error {
	err := _blockService.IBlockRepo.CreateBlock(block)
	return err
}

