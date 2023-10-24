package postgres

import (
	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDatabase(config *config.EnvConfig) (*gorm.DB, error) {
	//create db instance wuth gorm
	db, err := gorm.Open(postgres.Open(config.DbUrl))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&entities.UserEntity{})
	return db, nil
}
