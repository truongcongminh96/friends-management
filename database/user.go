package database

import "github.com/friends-management/models"

func (db Database) GetUserList() (*models.UserListResponse, error) {
	list := &models.UserListResponse{}
	rows, err := db.Conn.Query("SELECT * FROM users")
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Email)
		if err != nil {
			return list, err
		}
		list.Users = append(list.Users, user)
	}
	return list, nil
}

func (db Database) CheckUserExist(email string) (bool, error) {
	var count int
	query := `
	SELECT count(*) 
	FROM users u 
	WHERE 
		u.email = $1;`
	row := db.Conn.QueryRow(query, email)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	}
	return true, nil
}

func (db Database) CreateUser(email string) error {
	query := `INSERT INTO users (email) VALUES ($1);`
	_, err := db.Conn.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) CreateFriendConnection(friends []string) error {
	emailUserOne := friends[0]
	emailUserTwo := friends[1]

	query := `INSERT INTO friend (emailuserone, emailusertwo) VALUES ($1, $2);`
	_, err := db.Conn.Exec(query, emailUserOne, emailUserTwo)

	if err != nil {
		return err
	}
	return nil
}

func (db Database) RetrieveFriendListByEmail(email string) (*models.FriendListResponse, error) {
	friendList := &models.FriendListResponse{}
	query := `
	SELECT emailuserone, id 
	FROM friend  
	WHERE 
		emailusertwo = $1
	UNION 
	SELECT emailusertwo, id 
	FROM friend 
	WHERE 
		emailuserone =  $1 
	ORDER BY id;`
	rows, err := db.Conn.Query(query, email)
	if err != nil {
		return friendList, err
	}
	var id int
	for rows.Next() {
		var item models.User
		err := rows.Scan(&item.Email, &id)
		if err != nil {
			return friendList, err
		}
		friendList.Friends = append(friendList.Friends, item.Email)
	}
	return friendList, nil
}

func (db Database) CreateSubscribe(requestor, target string) error {
	query := `INSERT INTO subscription (requestor, target) VALUES ($1, $2);`
	_, err := db.Conn.Exec(query, requestor, target)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) CreateBlockFriend(requestor, target string) error {
	query := `INSERT INTO block (requestor, target) VALUES ($1, $2);`
	_, err := db.Conn.Exec(query, requestor, target)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) GetAllBlockerByEmail(requestor string) (*models.User, error) {
	targetList := &models.User{}
	query := `SELECT b.requestor FROM block b WHERE b.target = $1;`
	rows, err := db.Conn.Query(query, requestor)
	if err != nil {
		return targetList, err
	}
	for rows.Next() {
		var item models.BlockRequest
		err := rows.Scan(&item.Requestor)
		if err != nil {
			return targetList, err
		}
		targetList.Blocked = append(targetList.Blocked, item.Requestor)
	}
	return targetList, nil
}

func (db Database) GetAllSubscriber(requestor string) (*models.User, error) {
	targetList := &models.User{}
	query := `SELECT s.requestor FROM subscription s WHERE s.target = $1;`
	rows, err := db.Conn.Query(query, requestor)
	if err != nil {
		return targetList, err
	}
	for rows.Next() {
		var item models.SubscriptionRequest
		err := rows.Scan(&item.Requestor)
		if err != nil {
			return targetList, err
		}
		targetList.Subscription = append(targetList.Subscription, item.Requestor)
	}
	return targetList, nil
}

func (db Database) CheckIsFriend(email []string) (bool, error) {
	var countUserOne int
	var countUserTwo int

	query := `
	SELECT COUNT(*) 
		FROM friend  
	WHERE 
		emailuserone = $1 AND emailusertwo = $2;`
	row := db.Conn.QueryRow(query, email[0], email[1])
	err := row.Scan(&countUserOne)
	if err != nil {
		return false, err
	}

	query = `SELECT COUNT(*) 
		FROM friend  
	WHERE 
		emailuserone = $2 AND emailusertwo = $1;`

	row = db.Conn.QueryRow(query, email[0], email[1])
	err = row.Scan(&countUserTwo)

	if err != nil {
		return false, err
	}

	if countUserOne > 0 || countUserTwo > 0 {
		return true, nil
	}

	return false, nil
}
