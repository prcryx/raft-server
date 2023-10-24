package entities

type UserEntity struct {
	UserId     uint   `json:"UserId" gorm:"primaryKey; autoIncrement"`
	FirebaseId string `json:"FirebaseId" gorm:"<-:create"`
	Email      string `json:"Email" gorm:"unique"`
	CreatedAt  int64  `json:"CreatedAt" gorm:"autoUpdateTime:nano"`
}
