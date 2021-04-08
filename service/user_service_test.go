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
			friend: []string{"andy@example", "lisa@example"},
			expectedResult: &models.ResultResponse{
				Success: false,
			},
		},
	}
	data := DbInstance{db}
	err := insertConnectFriend(db.Conn)
	assert.NoError(t, err)
	for _, tt := range testCases {
		req := &models.FriendConnectionRequest{
			Friends: tt.friend,
		}
		response, err := data.CreateFriendConnection(req.Friends)
		assert.NoError(t, err)
		assert.Equal(t, tt.expectedResult, response)
	}
}

func insertConnectFriend(db *sql.DB) error {
	query :=
		`
		truncate friend;
		truncate users cascade;
		insert into users (email) values ('andy@example');
		insert into users (email) values ('john@example');
		insert into users (email) values ('ana@example');
		insert into users (email) values ('tom@example');
		insert into users (email) values ('jerry@example');
		insert into block (requestor, target) values ('tom@example','jerry@example');
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

