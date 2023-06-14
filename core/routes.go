package core

// initialises Will register all routes
func (s *Server) initaliseRoutes() {
	s.router.HandleFunc(pfProxyPath, s.reverseProxyHandler())
	s.router.HandleFunc(pfLoginUrl, s.loginHandler())
	s.router.HandleFunc(pfLogoutUrl, s.logoutHandler())
	s.router.HandleFunc(pfStatusUrl, s.statusHandler())
	s.router.HandleFunc(pfCallbackUrl, s.callBack())
}
