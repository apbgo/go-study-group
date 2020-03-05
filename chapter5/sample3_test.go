package chapter5

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser3_UserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockIFDBService2(ctrl)

	expect := "User1"
	mockService.EXPECT().Get(1).Return(UserData{
		Id:       1,
		UserName: expect,
	})
	user := User3{
		dbService: mockService,
	}
	assert.Equal(t, expect, user.UserName(1))
}

func TestUser3_Mock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockIFDBService2(ctrl)

	// 呼び出される順番を指定
	gomock.InOrder(
		mockService.EXPECT().Get(1).Return(UserData{Id: 1, UserName: "UserA"}),
		mockService.EXPECT().Get(2).Return(UserData{Id: 1, UserName: "UserB"}),
	)

	mockService.EXPECT().Get(1).Return(UserData{Id: 1, UserName: "UserA"}).Times(1)
	mockService.EXPECT().Get(1).Return(UserData{Id: 1, UserName: "UserA"}).AnyTimes()
	mockService.EXPECT().Get(1).Return(UserData{Id: 1, UserName: "UserA"}).MinTimes(1)
	mockService.EXPECT().Get(1).Return(UserData{Id: 1, UserName: "UserA"}).MaxTimes(1)

}
