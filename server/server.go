package server

// Init gin server
func Init() {
	r := NewRouter()
	r.Run("8080")
}
