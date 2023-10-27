package entities

type UserEntity struct {
	UID       uint   `json:"uid" gorm:"primaryKey; autoIncrement"`
	PhoneNo   string `json:"phoneNo"`
	CreatedAt int64  `json:"createdAt" gorm:"autoUpdateTime:nano"`
}