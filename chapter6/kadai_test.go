package chapter6

import (
	"context"
	"database/sql"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// このテストはgomockを利用するサンプルです。
func TestSample(t *testing.T) {
	t.Run("サンプル1", func(t *testing.T) {
		// サブテストごとにMockのControllerを作成してください。
		ctrl := gomock.NewController(t)

		// 自動生成されたMockをNewする
		mock := NewMockIFUserItemRepository(ctrl)

		// ここからは意味がないテスト
		ctx := context.Background()
		userItem := IUserItem{
			UserID: 1,
			ItemID: 1,
			Count:  100,
		}

		mock.EXPECT().
			// ここに渡された変数は、値が一致しない場合はテストが成功しない（ポインタの場合はポインタの一致）
			// gomock.Any()は全ての値が許容される
			// *sql.TxやiUserItemはService内で生成されるため、ポインタの一致をチェックすることは難しい
			Update(ctx, gomock.Any(), gomock.Any()).
			// そのためDoAndReturnの関数内で値をチェックしてあげるとよい
			DoAndReturn(
				func(ctx1 context.Context, tx1 *sql.Tx, userItem1 *IUserItem) (ok bool, err error) {
					assert.Equal(t, userItem, *userItem1)
					// この関数の戻り値がMockを実行した時の戻り値になる
					return true, nil
				},
			)

		ok, err := mock.Update(ctx, nil, &userItem)
		assert.True(t, ok)
		assert.NoError(t, err)
	})
}

func TestUserItemService_Provide(t *testing.T) {
	t.Skip("[実装対象]")
}
