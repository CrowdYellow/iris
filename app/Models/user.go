package Models

import "time"

type User struct {
	Id        int64
	Name      string    `xorm:"varchar(200)"`
	Phone     string    `xorm:"varchar(12)"`
	Password  string    `xorm:"varchar(200)"`
	RoleId    int64     `xorm:"int(11)"`
	Enable    bool      `xorm:"int(11)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
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
