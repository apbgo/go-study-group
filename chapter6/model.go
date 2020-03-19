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
	DeletedAt         sql.NullTime   `xorm:"'deleted_at' deleted"`
}

type IUserItem struct {
	UserID    int64        `xorm:"'user_id' pk"`
	ItemID    int64        `xorm:"'item_id'"`
	Count     int64        `xorm:"'count'"`
	CreatedAt time.Time    `xorm:"'created_at' created"`
	UpdatedAt time.Time    `xorm:"'updated_at' updated"`
	DeletedAt sql.NullTime `xorm:"'deleted_at' deleted"`
}

type JoinedUser struct {
	IUser     `xorm:"extends"`
	IUserItem `xorm:"extends"`
}
