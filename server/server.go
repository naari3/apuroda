package server

// Start Start
func Start() {
	r := newRouter()

	r.Run()
}
