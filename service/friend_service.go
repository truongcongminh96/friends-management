package service

import (
	"github.com/friends-management/helper"
	"github.com/friends-management/models"
	"github.com/friends-management/repositories"
)

type IFriendService interface {
	CreateFriend(friend *models.Friend) error
	CheckExistedFriend(userId1 int, userId2 int) (bool, error)
	CheckBlockedByUser(requestorId int, targetId int) (bool, string, error)
	GetFriendsList(userId int) ([]string, error)
	GetCommonFriends(Ids []int) ([]string, error)
	GetEmailsReceiveUpdate(senderId int, text string) ([]string, error)
}

type FriendService struct {
	IFriendRepo repositories.IFriendRepo
	IUserRepo   repositories.IUserRepo
}

func (_friendService FriendService) CreateFriend(friend *models.Friend) error {
	err := _friendService.IFriendRepo.CreateFriend(friend)
	return err
}

func (_friendService FriendService) CheckExistedFriend(userId1 int, userId2 int) (bool, error) {
	isExist, err := _friendService.IFriendRepo.IsExistedFriend(userId1, userId2)
	return isExist, err
}

func (_friendService FriendService) GetFriendsList(userId int) ([]string, error) {
	listFriendId, err := _friendService.IFriendRepo.GetListFriendId(userId)
	if err != nil {
		return nil, err
	}

	// Get users that is blocked by this user and get users that block this user
	blockMap := make(map[int]bool)
	blockedByTargetIds, err := _friendService.IFriendRepo.GetIdsBlockedByTarget(userId)
	if err != nil {
		return nil, err
	}

	for _, id := range blockedByTargetIds {
		blockMap[id] = true
	}

	blockedByUserIds, err := _friendService.IFriendRepo.GetIdsBlockedUsers(userId)
	if err != nil {
		return nil, err
	}
	for _, id := range blockedByUserIds {
		blockMap[id] = true
	}

	// Get friends that has friend connection and not blocking
	friendIds := make([]int, 0)
	for _, id := range listFriendId {
		if _, ok := blockMap[id]; !ok {
			friendIds = append(friendIds, id)
		}
	}

	result, err := _friendService.IUserRepo.GetEmailsByIDs(friendIds)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (_friendService FriendService) CheckBlockedByUser(requestorId int, targetId int) (bool, string, error) {
	isBlocked, message, err := _friendService.IFriendRepo.IsBlockedByUser(requestorId, targetId)
	return isBlocked, message, err
}

func (_friendService FriendService) GetCommonFriends(Ids []int) ([]string, error) {
	// Get friends list of each email address
	friends1, err := _friendService.GetFriendsList(Ids[0])
	if err != nil {
		return nil, err
	}
	friends2, err := _friendService.GetFriendsList(Ids[1])
	if err != nil {
		return nil, err
	}

	// Find common friends in 2 friends list
	var commonFriends []string
	hashMap := make(map[string]bool)
	for _, email := range friends1 {
		hashMap[email] = true

	}
	for _, email := range friends2 {
		if _, ok := hashMap[email]; ok {
			commonFriends = append(commonFriends, email)
		}
	}
	return commonFriends, nil
}

func (_friendService FriendService) GetEmailsReceiveUpdate(senderId int, text string) ([]string, error) {
	// Get emails that blocked from the sender
	blockedIds, err := _friendService.IFriendRepo.GetIdsBlockedUsers(senderId)
	if err != nil {
		return nil, err
	}

	blockedEmails, err := _friendService.IUserRepo.GetEmailsByIDs(blockedIds)
	if err != nil {
		return nil, err
	}

	var blockedMap = make(map[string]bool)
	for _, email := range blockedEmails {
		blockedMap[email] = true
	}

	result := make([]string, 0)

	emailMap := make(map[string]bool)

	// Get emails that has a friend connection with the sender
	friendIds, err := _friendService.IFriendRepo.GetListFriendId(senderId)
	if err != nil {
		return nil, err
	}

	friendEmails, err := _friendService.IUserRepo.GetEmailsByIDs(friendIds)
	if err != nil {
		return nil, err
	}

	for _, email := range friendEmails {
		// If not blocked
		if _, ok := blockedMap[email]; !ok {
			// Append to result and add to emailMap
			result = append(result, email)
			emailMap[email] = true
		}
	}

	// Get emails that subscribe to updates from the sender
	subscriberIds, err := _friendService.IFriendRepo.GetIdsSubscribers(senderId)
	if err != nil {
		return nil, err
	}

	subscriberEmails, err := _friendService.IUserRepo.GetEmailsByIDs(subscriberIds)
	if err != nil {
		return nil, err
	}

	for _, email := range subscriberEmails {
		// If not blocked
		if _, ok := blockedMap[email]; !ok {
			// If not in emailMap then append to result and add to map
			if _, ok := emailMap[email]; !ok {
				result = append(result, email)
				emailMap[email] = true
			}
		}
	}

	//Add mentionedEmails to result
	mentionedEmails := helper.FindEmailFromText(text)
	for _, email := range mentionedEmails {
		// If not blocked
		if _, ok := blockedMap[email]; !ok {
			// If not in emailMap then append to result and add to map
			if _, ok := emailMap[email]; !ok {
				result = append(result, email)
				emailMap[email] = true
			}
		}
	}

	return result, nil
}
