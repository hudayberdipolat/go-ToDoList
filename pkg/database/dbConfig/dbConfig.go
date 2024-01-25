package dbConfig

import (
	"fmt"
	"github.com/hudayberdipolat/go-ToDoList/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	config *config.Config
}

func NewDBConnection(config *config.Config) DbConfig {
	return DbConfig{
		config: config,
	}
}

func (dbConfig *DbConfig) GetDBConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.config.DbConfig.DbHost,
		dbConfig.config.DbConfig.DbUser,
		dbConfig.config.DbConfig.DbPassword,
		dbConfig.config.DbConfig.DbName,
		dbConfig.config.DbConfig.DbPort,
		dbConfig.config.DbConfig.DbSllMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil

}
