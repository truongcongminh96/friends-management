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
	SELECT emailusertwo, id 
	FROM friend  
	WHERE 
		emailuserone = $1 
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
