package main

import (
	"github.com/RolandKun5/go-rest-api/src/database"
	"github.com/RolandKun5/go-rest-api/src/server"
)

func main() {
	database.Init()
	server.Start()
}
