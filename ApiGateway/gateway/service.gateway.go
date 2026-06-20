package gateway

import (
	"ApiGateway/models"
)

type ServiceRegistry struct {
	services map[string]models.Service
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string]models.Service),
	}
}

func (s *ServiceRegistry) Register(prefix string, service models.Service) {
	s.services[prefix] = service
}

func (s *ServiceRegistry) Resolve(path string) (*models.Service, bool) {
	service, ok := s.services[path]
	if !ok {
		return nil, ok
	}
	return &service, ok
}

// RegisterAll registers all microservices from config
func (s *ServiceRegistry) RegisterAll(authServiceURL, uploadServiceURL string) {
	if authServiceURL != "" {
		s.Register("/auth", models.Service{
			Name: "AuthenticationService",
			URL:  authServiceURL,
		})
	}
	if uploadServiceURL != "" {
		s.Register("/upload", models.Service{
			Name: "UploadService",
			URL:  uploadServiceURL,
		})
	}
}
