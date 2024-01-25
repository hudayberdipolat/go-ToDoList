package app

import (
	"github.com/hudayberdipolat/go-ToDoList/pkg/config"
	"github.com/hudayberdipolat/go-ToDoList/pkg/database/dbConfig"
	CustomHttp "github.com/hudayberdipolat/go-ToDoList/pkg/http"
	"gorm.io/gorm"
	"net/http"
)

type Dependencies struct {
	Config     *config.Config
	DB         *gorm.DB
	HttpClient *http.Client
}

func GetDependencies() (*Dependencies, error) {
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	// db config
	newDBConfig := dbConfig.NewDBConnection(getConfig)
	db, errDB := newDBConfig.GetDBConnection()
	if errDB != nil {
		return nil, errDB
	}
	// http
	httpClient := CustomHttp.NewHttpClient()
	return &Dependencies{
		Config:     getConfig,
		DB:         db,
		HttpClient: httpClient,
	}, nil
}
