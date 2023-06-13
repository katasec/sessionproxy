package core

// initialises Will register all routes
func (s *Server) initaliseRoutes() {
	s.router.HandleFunc("/", s.reverseProxyHandler())
	s.router.HandleFunc("/.pathfinder/login", s.loginHandler())
	s.router.HandleFunc("/.pathfinder/logout", s.logoutHandler())
	s.router.HandleFunc("/.pathfinder/status", s.statusHandler())
	s.router.HandleFunc("/.pathfinder/callback", s.callBack())
}
