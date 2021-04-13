package database_test

import (
	"github.com/friends-management/database"
	_ "github.com/friends-management/database"
	_ "github.com/lib/pq"
	"log"
	"testing"
)

type TestCase struct {
	excepted bool
	actual   bool
}

func TestConnectDB(t *testing.T) {
	t.Run("return if connected", func(t *testing.T) {
		testCase := TestCase{
			excepted: true,
			actual:   false,
		}

		_, err := database.ConnectDB("postgres", "1", "friends-management")

		if err == nil {
			testCase.actual = true
		}

		if testCase.actual != testCase.excepted {
			t.Fail()
			log.Println(err.Error())
		}
	})
}

func TestFailConnectDB(t *testing.T) {
	t.Run("return if not connected", func(t *testing.T) {
		testCase := TestCase{
			excepted: false,
			actual:   false,
		}

		_, err := database.ConnectDB("postgres", "111", "friends-management")

		if err == nil {
			testCase.actual = true
		}

		if testCase.actual != testCase.excepted {
			t.Fail()
			log.Println(err.Error())
		}
	})
}
