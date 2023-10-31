package postgres

import (
	"time"

	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/internal/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreatePostgresDatabase(config *config.EnvConfig) (*gorm.DB, error) {
	//create db instance wuth gorm
	db, err := gorm.Open(postgres.Open(config.DB.URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Get generic database object sql.DB to use its functions
	//use dbResolver in future!!
	// db.Use()

	if sqlDB, err := db.DB(); err != nil {
		return nil, err
	} else {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Hour)
	}

	//migrate the entities here
	db.AutoMigrate(
		&entities.UserEntity{},
	)

	return db, nil
}
