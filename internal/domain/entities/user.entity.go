package entities

type UserEntity struct {
	UID       uint   `json:"uid" gorm:"primaryKey; autoIncrement; column: uid"`
	PhoneNo   string `json:"phoneNo" gorm:"column:phoneNo"`
	CreatedAt int64  `json:"createdAt" gorm:"autoUpdateTime:nano; column:createdAt"`
}

type User struct {
	*UserEntity
	AccessToken string `json:"accessToken"`
}

func (userEntity *UserEntity) ToUser(accessToken string) *User {
	return &User{
		UserEntity:  userEntity,
		AccessToken: accessToken,
	}
}
