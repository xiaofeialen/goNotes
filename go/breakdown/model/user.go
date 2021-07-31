package model

type User struct {
	ID       int    `gorm:"column:id;primary_key;auto"`
	Username string `gorm:"column:username;size(32)"`
	Password string `gorm:"column:password;size(128)"`
	Name     string `gorm:"column:name;size(128);null"`
}

func (t *User) TableName() string {
	return "user"
}

func GetUserByName(name string) (user User, err error) {
	err = db.Where("username = ?", name).First(&user).Error
	return
}

func UserInfoToFilter() (users []User, err error) {
	err = db.Model(&User{}).Find(&users).Error
	return
}
