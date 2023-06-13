package core

// initialises Will register all routes
func (s *Server) initaliseRoutes() {
	s.router.HandleFunc("/", s.helloHandlerFunc())
	//s.router.HandleFunc("/", s.helloHandler)
	s.router.HandleFunc("/hello", s.helloHandlerFunc())
	s.router.HandleFunc("/.pathfinder/login", s.loginHandler())
	s.router.HandleFunc("/.pathfinder/logout", s.logoutHandler())
	s.router.HandleFunc("/.pathfinder/status", s.helloHandler)
	s.router.HandleFunc("/.pathfinder/callback", s.callBack())
}
