package core

// initialises Will register all routes
func (s *Server) initaliseRoutes() {
	s.helloRoute()
}

// helloRoute is a test route
func (s *Server) helloRoute() {
	s.router.HandleFunc("/hello", s.helloHandler())
}
