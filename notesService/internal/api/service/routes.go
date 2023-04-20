package service

func (s *Service) initRoutes() {
	s.S.GET("/note", s.Get)
	s.S.POST("/add", s.Add)
	s.S.DELETE("/delete", s.Delete)
}
