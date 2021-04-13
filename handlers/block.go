package handlers

import (
	"encoding/json"
	"errors"
	"github.com/friends-management/models"
	"github.com/friends-management/service"
	"github.com/go-chi/render"
	"net/http"
)

type BlockHandlers struct {
	IBlockService service.IBlockService
	IUserService  service.IUserService
}

func (_blockHandlers *BlockHandlers) CreateBlock(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var blockRequest models.BlockRequest
	if err := json.NewDecoder(r.Body).Decode(&blockRequest); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// Validate request
	if err := blockRequest.Validate(); err != nil {
		_ = render.Render(w, r, ErrBadRequest)
		return
	}

	// handlers block
	Ids, statusCode, err := _blockHandlers.checkBlock(blockRequest)
	if err != nil {
		_ = render.Render(w, r, ErrorRenderer(err, statusCode))
		return
	}

	blockModel := &models.Block{
		Requestor: Ids[0],
		Target:    Ids[1],
	}
	// Call service to create blocking
	if err := _blockHandlers.IBlockService.CreateBlock(blockModel); err != nil {
		_ = render.Render(w, r, ServerErrorRenderer(err))
		return
	}

	// Response
	err = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
	})
	if err != nil {
		return
	}
	return
}

func (_blockHandlers *BlockHandlers) checkBlock(block models.BlockRequest) ([]int, int, error) {
	// Get user id of the requestor
	requestorID, err := _blockHandlers.IUserService.GetUserIDByEmail(block.Requestor)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if requestorID == 0 {
		return nil, http.StatusBadRequest, errors.New("the requestor does not exist")
	}

	// Get user id of the target
	targetID, err := _blockHandlers.IUserService.GetUserIDByEmail(block.Target)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if targetID == 0 {
		return nil, http.StatusBadRequest, errors.New("the target does not exist")
	}

	// Check if blocking exists
	exist, err := _blockHandlers.IBlockService.CheckExistedBlock(requestorID, targetID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if exist {
		return nil, http.StatusAlreadyReported, errors.New("you are block the target")
	}
	return []int{requestorID, targetID}, 0, nil
}
