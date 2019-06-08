package Models

import (
	"log"
	"time"
)

type User struct {
	Id        int64
	Name      string    `xorm:"varchar(200) notnull unique" json:"name" form:"name"`
	NickName  string    `xorm:"varchar(200) notnull" json:"name" form:"name"`
	Avatar    string    `xorm:"varchar(200) notnull"`
	Phone     string    `xorm:"varchar(12) notnull unique" json:"phone" form:"phone"`
	Password  string    `xorm:"varchar(200) notnull" json:"password" form:"password"`
	RoleId    int64     `xorm:"int(11) notnull"`
	Enable    bool      `xorm:"int(11) notnull"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func init() {
	err := DB.Sync2(new(User))

	if err != nil {
		log.Print("用户表创建失败", err)
	}
}

// 创建用户
func CreateUser(user ...*User) (int64, error) {
	return DB.Insert(user)
}

// 根据用户名获取用户
func GetUserByModelsUser(user *User) (bool, error) {
	return DB.Get(user)
}

// 获取所有用户
func GetUserList() ([]*User) {
	userList := make([]*User, 0)
	_ = DB.Find(&userList)
	return userList
}

// 根据用户ID编辑用户
func UpdateUserById(user *User) (int64, error) {
	return DB.Id(user.Id).Update(user)
}

// 根据ID删除用户
func DeleteUserById(uIds []int64) (effect int64, err error) {
	user := new(User)
	for _, v := range uIds {
		i, err1 := DB.Id(v).Delete(user)
		effect += i
		err = err1
	}
	return
}
