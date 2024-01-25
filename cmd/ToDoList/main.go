package main

import (
	"fmt"
	"github.com/hudayberdipolat/go-ToDoList/internal/app"
	"github.com/hudayberdipolat/go-ToDoList/internal/setup/constructor"
	"log"
)

func main() {
	// get dependencies
	appDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal("app dependencies error :-->", err.Error())
	}
	// build constructor
	constructor.Build(appDependencies)
	// app router
	appRouter := app.NewApp(appDependencies)
	runServer := fmt.Sprintf("%s:%s",
		appDependencies.Config.HttpServer.ServerHost, appDependencies.Config.HttpServer.ServerPort)
	if errRunServer := appRouter.Listen(runServer); errRunServer != nil {
		log.Fatal("error run server : --->", errRunServer.Error())
	}
}
