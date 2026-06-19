package gateway

import "ApiGateway/models"

type ServiceRegistry struct{
	routes map[string]models.Service
}

func NewServiceRegistry() *ServiceRegistry{
	return &ServiceRegistry{}
}

func (s *ServiceRegistry) Register(prefix string, service models.Service){
	s.routes[prefix]=service
}

func (s *ServiceRegistry) Resolve(path string) (models.Service, bool) {
	service, ok:=s.routes[path]
	if !ok{
		return models.Service{}, ok
	}
	return service, ok
}