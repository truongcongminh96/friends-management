package service

import (
	"database/sql"
	"fmt"
	"github.com/friends-management/database"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFriendConnection(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()
	testCases := []struct {
		name           string
		friend         []string
		expectedResult *models.ResultResponse
	}{
		{
			name:   "Success",
			friend: []string{"andy@example", "john@example"},
			expectedResult: &models.ResultResponse{
				Success: true,
			},
		},
		{
			name:   "duplicate connect friends",
			friend: []string{"andy@example", "john@example"},
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
		{
			name:   "connect failed by email address does not exist",
			friend: []string{"andy@example", "lasi@example"},
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
	}
	data := DbInstance{db}
	err := insertConnectFriend(db.Conn)
	assert.NoError(t, err)
	for _, tc := range testCases {
		req := &models.FriendConnectionRequest{
			Friends: tc.friend,
		}
		response, err := data.CreateFriendConnection(req.Friends)
		assert.NoError(t, err)
		assert.Equal(t, tc.expectedResult, response)
	}
}

func insertConnectFriend(db *sql.DB) error {
	query :=
		`
		truncate friend;
		truncate users cascade;
		insert into users (email) values ('andy@example');
		insert into users (email) values ('john@example');
		insert into users (email) values ('lisa@example');
		insert into users (email) values ('kate@example');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func TestRetrieveFriendList(t *testing.T) {
	db := createConnectionForTest()
	defer db.Conn.Close()

	testCases := struct {
		name           string
		email          string
		expectedResult *models.FriendListResponse
	}{

		name:  "success",
		email: "andy@example",
		expectedResult: &models.FriendListResponse{
			Success: true,
			Friends: []string{"john@example", "lisa@example"},
			Count:   2,
		},
	}

	data := DbInstance{db}
	err := insertFriend(db.Conn)
	assert.NoError(t, err)
	req := &models.FriendListRequest{
		Email: testCases.email,
	}
	response, err := data.RetrieveFriendList(req.Email)
	assert.NoError(t, err)
	assert.Equal(t, testCases.expectedResult, response)
}

func insertFriend(db *sql.DB) error {
	query :=
		`
		truncate friend;
		truncate users cascade;
		insert into users (email) values ('andy@example');
		insert into users (email) values ('john@example');
		insert into users (email) values ('lisa@example');
		insert into users (email) values ('kate@example');
		insert into friend (emailuserone, emailusertwo) values ('andy@example','john@example');
		insert into friend (emailuserone, emailusertwo) values ('andy@example','lisa@example');
		insert into friend (emailuserone, emailusertwo) values ('lisa@example','kate@example');
		`
	_, err := db.Exec(query)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func createConnectionForTest() database.Database {
	db := database.Database{}
	// Open the connection
	conn, err := sql.Open("postgres", "postgres://postgres:1@localhost:5432/friends-management?sslmode=disable")

	if err != nil {
		panic(err)
	}
	db.Conn = conn
	// check the connection
	err = db.Conn.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
