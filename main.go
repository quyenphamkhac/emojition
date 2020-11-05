package main

import (
	"github.com/quyenphamkhac/emojition/db"
	"github.com/quyenphamkhac/emojition/server"
)

func main() {
	db.Init()
	server.Init()
}
