package models

import (
	"errors"
	"strings"
)

type Block struct {
	Requestor int
	Target    int
}

type BlockRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func (_blockRequest BlockRequest) Validate() error {
	var errMessages []string
	if _blockRequest.Requestor == "" {
		errMessages = append(errMessages, "requestor is required")
	}
	if _blockRequest.Target == "" {
		errMessages = append(errMessages, "target is required")
	}
	if len(errMessages) > 0 {
		return errors.New(strings.Join(errMessages, "\n"))
	}
	return nil
}
