package chapter5

//go:generate mockgen -package chapter5 -destination sample3_mock.go -self_package=github.com/apbgo/go-study-group/chapter5 github.com/apbgo/go-study-group/chapter5 IFDBService2
type IFDBService2 interface {
	Get(id int) UserData
}

type User3 struct {
	dbService IFDBService
}

func (u *User3) UserName(id int) string {
	user := u.dbService.Get(id)
	return user.UserName
}
