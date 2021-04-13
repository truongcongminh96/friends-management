package models

import (
	"errors"
	"strings"
)

type Subscribe struct {
	Requestor int
	Target    int
}

type SubscribeRequest struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

func (subscribeRequest SubscribeRequest) Validate() error {
	var errMessages []string
	if subscribeRequest.Requestor == "" {
		errMessages = append(errMessages, "requestor is required")
	}
	if subscribeRequest.Target == "" {
		errMessages = append(errMessages, "target is required")
	}
	if len(errMessages) > 0 {
		return errors.New(strings.Join(errMessages, "\n"))
	}
	return nil
}
