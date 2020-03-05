package chapter5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser2_UserName(t *testing.T) {
	user := User2{
		dbService: TestService{},
	}
	assert.Equal(t, "UserA", user.UserName(1))
}

// 差し替える用のstruct
type TestService struct{}

func (t TestService) Get(id int) UserData {
	return UserData{
		Id:       1,
		UserName: "UserA",
	}
}
