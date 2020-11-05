package server

import "os"

// Init gin server
func Init() {
	r := NewRouter()
	r.Run(os.Getenv("HOST"))
}
