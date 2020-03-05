package chapter5

type UserData struct {
	Id       int
	UserName string
}

// before ------------------------------------------
type MySQLService struct{}

func (s MySQLService) Get(id int) UserData {
	// MySQLに接続してUserDataを取得する処理
	// 取得したUserDataを返却
	return UserData{}
}

type User struct {
	dbService MySQLService
}

func (u User) UserName(id int) string {
	user := u.dbService.Get(id)
	return user.UserName
}

// after ------------------------------------------
type IFDBService interface {
	Get(id int) UserData
}

type User2 struct {
	dbService IFDBService
}

func (u *User2) UserName(id int) string {
	user := u.dbService.Get(id)
	return user.UserName
}
