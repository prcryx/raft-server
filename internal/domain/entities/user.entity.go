package entities

type UserEntity struct {
	UserId      uint   `json:"UserId" gorm:"primaryKey; autoIncrement"`
	FirebaseId  string `json:"FirebaseId"`
	DisplayName string `json:"DisplayName"`
	Email       string `json:"Email"`
}
