package chapter6

import (
	"database/sql"
	"time"
)

type IUser struct {
	UserID            int64          `xorm:"'user_id' pk"`
	OSType            int64          `xorm:"'os_type'"`
	Name              sql.NullString `xorm:"'name'"`
	GamestartDatetime time.Time      `xorm:"'gamestart_datetime'"`
	LatestVersion     int64          `xorm:"'latest_version'"`
	CreatedAt         time.Time      `xorm:"'created_at' created"`
	UpdatedAt         time.Time      `xorm:"'updated_at' updated"`
}

type IUserStatus struct {
	UserID    int64     `xorm:"'user_id' pk"`
	Level     int64     `xorm:"'level'"`
	Stumina   int64     `xorm:"'stumina'"`
	EXP       int64     `xorm:"'exp'"`
	CreatedAt time.Time `xorm:"'created_at' created"`
	UpdatedAt time.Time `xorm:"'updated_at' updated"`
}

type JoinedUser struct {
	IUser       `xorm:"extends"`
	IUserStatus `xorm:"extends"`
}
