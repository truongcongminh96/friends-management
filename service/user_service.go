package service

import (
	"github.com/friends-management/database"
	"github.com/friends-management/models"
	"log"
)

type DbInstance struct {
	Db database.Database
}

type IUserService interface {
	GetUserList() (*models.UserListResponse, error)
	CreateUser(email string) (*models.ResultResponse, error)
	CreateFriendConnection(req []string) (*models.ResultResponse, error)
	RetrieveFriendList(email string) (*models.FriendListResponse, error)
	GetCommonFriendsList(req []string) (*models.FriendListResponse, error)
	CreateSubscribe(req *models.SubscriptionRequest) (*models.ResultResponse, error)
	CreateBlockFriend(req *models.BlockRequest) (*models.ResultResponse, error)
	CreateReceiveUpdate(emailSender string) (*models.SendUpdateEmailResponse, error)
}

func (db DbInstance) GetUserList() (*models.UserListResponse, error) {
	response := &models.UserListResponse{}

	rep, err := db.Db.GetUserList()

	if err != nil {
		return response, err
	}

	response.Users = rep.Users
	return response, nil
}

func (db DbInstance) CreateUser(email string) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}
	if err := db.Db.CreateUser(email); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (db DbInstance) CreateFriendConnection(req []string) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}

	isUserExist, err := db.Db.CheckUserExist(req[0])
	if !isUserExist {
		log.Printf("Your email address does not register")
		response.Success = false
		return response, err
	}

	isYourFriendExist, err := db.Db.CheckUserExist(req[1])
	if !isYourFriendExist {
		log.Printf("Email address your friend does not register")
		response.Success = false
		return response, err
	}

	isFriend , err := db.Db.CheckIsFriend(req)

	if err != nil {
		return response, err
	}

	if isFriend {
		log.Printf("You are now friend")
		response.Success = false
		return response, nil
	}

	if err := db.Db.CreateFriendConnection(req); err != nil {
		response.Success = false
		return response, err
	}
	response.Success = true
	return response, nil
}

func (db DbInstance) RetrieveFriendList(email string) (*models.FriendListResponse, error) {
	response := &models.FriendListResponse{}

	isUserExist, err := db.Db.CheckUserExist(email)
	if !isUserExist {
		log.Printf("Your email address does not register")
		response.Success = false
		return response, err
	}

	data, err := db.Db.RetrieveFriendListByEmail(email)
	if err != nil {
		return response, err
	}
	response.Success = true
	response.Friends = data.Friends
	response.Count = len(data.Friends)
	return response, nil
}

func (db DbInstance) GetCommonFriendsList(email []string) (*models.FriendListResponse, error) {
	response := &models.FriendListResponse{}

	isUserExist, err := db.Db.CheckUserExist(email[0])
	if !isUserExist {
		log.Printf("Your email address does not register")
		response.Success = false
		return response, err
	}

	isYourFriendExist, err := db.Db.CheckUserExist(email[1])
	if !isYourFriendExist {
		log.Printf("Email address your friend does not register")
		response.Success = false
		return response, err
	}

	friendListEmailOne, err := db.Db.RetrieveFriendListByEmail(email[0])
	if err != nil {
		return response, err
	}

	friendListEmailTwo, err := db.Db.RetrieveFriendListByEmail(email[1])
	if err != nil {
		return response, err
	}

	var mutualFriends []string

	// mutualFriends = append(friendListEmailOne.Friends, friendListEmailTwo.Friends...)
	// need to improve
	for _, friendOne := range friendListEmailOne.Friends {
		for _, friendTwo := range friendListEmailTwo.Friends {
			if friendOne == friendTwo {
				mutualFriends = append(mutualFriends, friendOne)
			}
		}
	}

	response.Success = true
	response.Friends = mutualFriends
	response.Count = len(mutualFriends)

	return response, nil
}

func (db DbInstance) CreateSubscribe(req *models.SubscriptionRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}

	isUserExist, err := db.Db.CheckUserExist(req.Requestor)
	if !isUserExist {
		log.Printf("Your email address does not register")
		response.Success = false
		return response, err
	}

	isYourSubscribeExist, err := db.Db.CheckUserExist(req.Target)
	if !isYourSubscribeExist {
		log.Printf("Email address you want to subcribe does not register")
		response.Success = false
		return response, err
	}

	isSubscribe, _ := db.Db.CheckSubscribe(req.Requestor, req.Target)

	if isSubscribe {
		log.Printf("You are subcriber your request")
		response.Success = false
		return response, err
	}

	if err := db.Db.CreateSubscribe(req.Requestor, req.Target); err != nil {
		return response, err
	}
	response.Success = true
	return response, nil
}

func (db DbInstance) CreateBlockFriend(req *models.BlockRequest) (*models.ResultResponse, error) {
	response := &models.ResultResponse{}

	isUserExist, err := db.Db.CheckUserExist(req.Requestor)
	if !isUserExist {
		log.Printf("Your email address does not register")
		response.Success = false
		return response, err
	}

	isYourBlockExist, err := db.Db.CheckUserExist(req.Target)
	if !isYourBlockExist {
		log.Printf("Email address you want to block does not register")
		response.Success = false
		return response, err
	}

	if err := db.Db.CreateBlockFriend(req.Requestor, req.Target); err != nil {
		return response, err
	}

	response.Success = true
	return response, nil
}

func (db DbInstance) CreateReceiveUpdate(emailSender string) (*models.SendUpdateEmailResponse, error) {
	response := &models.SendUpdateEmailResponse{}

	isUserExist, err := db.Db.CheckUserExist(emailSender)
	if !isUserExist {
		log.Printf("Your email address does not register")
		response.Success = false
		return response, err
	}

	listUsers, err := db.Db.GetUserList()

	blockedList, err := db.Db.GetAllBlockerByEmail(emailSender)
	if err != nil {
		return response, err
	}

	friendList, err := db.Db.RetrieveFriendListByEmail(emailSender)
	if err != nil {
		return response, err
	}

	subscriber, err := db.Db.GetAllSubscriber(emailSender)
	if err != nil {
		return response, err
	}

	var listUsersNotBlockSender []string
	for _, user := range listUsers.Users {
		for _, userBlock := range blockedList.Blocked {
			if user.Email != userBlock {
				listUsersNotBlockSender = append(listUsersNotBlockSender, user.Email)
			}
		}
	}

	var recipient []string
	if len(listUsersNotBlockSender) == 0 {
		recipient = append(recipient, friendList.Friends...)
	} else {
		var listFriendsNotBlock []string
		for _, userNotBlocked := range listUsersNotBlockSender {
			for _, friends := range friendList.Friends {
				if userNotBlocked == friends {
					listFriendsNotBlock = append(listFriendsNotBlock, userNotBlocked)
				}
			}
		}

		recipient = append(recipient, listFriendsNotBlock...)
	}

	recipient = append(recipient, subscriber.Subscription...)

	response.Success = true
	response.Recipients = recipient
	return response, nil
}
