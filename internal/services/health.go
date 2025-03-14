package services

import "backend-test/internal/interfaces"

type Healthcheck struct {
	HealthcheckRepository interfaces.IHealthcheckRepo
}

func (s *Healthcheck) HealthcheckServices() (string, error) {
	return "service healthy", nil
}
