package database

import "github.com/friends-management/models"

func (db Database) GetUserList() (*models.UserList, error) {
	list := &models.UserList{}
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
