package main

import (
	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
)

// go build -o app.exe ./cmd/main.go
// ./app.exe
func main() {
	dbFile := "scheduler.db"
	err := db.Init(dbFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	api.Init()
	err = server.LaunchServer()
	if err != nil {
		panic(err)
	}

}
