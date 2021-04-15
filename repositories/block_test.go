package repositories

import (
	"errors"
	"testing"

	"github.com/friends-management/helper"
	"github.com/friends-management/models"
	"github.com/stretchr/testify/require"
)

func TestBlockRepositories_CreateBlock(t *testing.T) {
	testCases := []struct {
		name          string
		input         *models.Block
		expectedError error
		fixturePath   string
	}{
		{
			name: "insert block failed",
			input: &models.Block{
				Requestor: 1,
				Target:    5,
			},
			expectedError: errors.New("pq: insert or update on table \"blocks\" violates foreign key constraint \"blocks_target_fkey\""),
			fixturePath:   "./test_migration/block/block.sql",
		},
		{
			name: "insert block successfully",
			input: &models.Block{
				Requestor: 3,
				Target:    4,
			},
			expectedError: nil,
			fixturePath:   "./test_migration/block/block.sql",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbMock := helper.ConnectDb()
			_ = helper.LoadFixture(dbMock, testCase.fixturePath)

			repo := BlockRepo{
				Db: dbMock,
			}

			err := repo.CreateBlock(testCase.input)

			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
