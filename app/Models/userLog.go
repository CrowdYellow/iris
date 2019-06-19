package Models

import (
	"log"
	"time"
)

type UserLog struct {
	Id        int64
	Uid       int64     `xorm:"int(11) notnull"`
	Do        string    `xorm:"varchar(200) notnull"`
	Device    string    `xorm:"varchar(200) notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func init() {
	err := DB.Sync2(new(UserLog))

	if err != nil {
		log.Print("用户记录表创建失败", err)
	}
}


//func CreateUserLog()  {
//
//}
