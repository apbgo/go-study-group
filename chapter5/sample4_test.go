package chapter5

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSample4_Method(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockIFCalcService(ctrl)

	gomock.InOrder(
		mockService.EXPECT().XXX(1).Return(4),
		mockService.EXPECT().XXX(4).Return(7),
		mockService.EXPECT().YYY(7, 2).Return(3),
	)
	calc := Calculator{
		service: mockService,
	}
	assert.Equal(t, 13, calc.Method(1, 4, 2))
}
