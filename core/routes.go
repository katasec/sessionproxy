package core

// initialises Will register all routes
func (s *Server) initaliseRoutes() {
	s.router.HandleFunc("/hello", s.helloHandler())
	s.router.HandleFunc("/login", s.loginHandler2())
	s.router.HandleFunc("/.pathfinder/status", s.callBack())
}
