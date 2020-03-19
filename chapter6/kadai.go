package chapter6

//go:generate mockgen -source=$GOFILE -destination=kadai_mock.go -package=$GOPACKAGE -self_package=github.com/apbgo/go-study-group/$GOPACKAGE

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// [課題内容]
// 以下の2つのInterface(IFUserItemService, IFUserItemRepository)を満たすstructを実装してください。
// 今回の課題ではトランザクション境界はProvide()内で構いません。
// gomockを利用してRepositoryのMockファイルを自動生成しています。
// テストではIFUserItemRepositoryをモックしたUserItemServiceを使ってみましょう。
// TransactionのBegin, Commit, Rollbackも本来Mockしたいですが、実装が多くなってしまうので
// 今回はする必要がありません。

// Reward 報酬モデル
type Reward struct {
	ItemID int64
	Count  int64
}

// IFUserItemService 報酬の付与の機能を表すインターフェイス
type IFUserItemService interface {
	// 対象のUserIDに引数で渡された報酬を付与します.
	Provide(ctx context.Context, userID int64, rewards ...Reward)
}

// IFUserItemRepository i_user_itemテーブルへの操作を行うインターフェイス
type IFUserItemRepository interface {
	// FindByUserIdAndItemIDs 一致するモデルを複数返却する.
	FindByUserIdAndItemIDs(
		ctx context.Context,
		tx *sql.Tx,
		userID int64,
		itemIDs []int64,
	) (iUserItems []*IUserItem, err error)

	// Insert 対象のモデルから1件Insertを実行する
	Insert(
		ctx context.Context,
		tx *sql.Tx,
		iUserItem *IUserItem,
	) error

	// Update対象のモデルから1件Updateを実行する
	// Update対象レコードが0件の場合、okはfalseになる
	Update(
		ctx context.Context,
		tx *sql.Tx,
		iUserItem *IUserItem,
	) (ok bool, err error)
}

// UserItemService [実装対象]
type UserItemService struct {
}

// NewUserItemService コンストラクタ [実装対象]
func NewUserItemService(
	db *sql.DB,
	userItemRepository *UserItemRepository,
) *UserItemService {
	return &UserItemService{}
}

// UserItemRepository [実装対象]
type UserItemRepository struct {
}

// NewUserItemRepository コンストラクタ [実装対象]
func NewUserItemRepository() *UserItemRepository {
	return &UserItemRepository{}
}
