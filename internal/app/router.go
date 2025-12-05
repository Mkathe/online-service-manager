package app

func (s *server) initRoutes() {
	s.app.Get("/healthz", s.HealthCheck)
	s.app.Get("/services", s.GetServices)
	s.app.Get("/services/total-cost", s.GetTotalCosts)
	s.app.Post("/services", s.CreateService)
	s.app.Put("/services/:id", s.UpdateService)
	s.app.Delete("/services/:id", s.DeleteService)
}
