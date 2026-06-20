package gateway

type RouteMapping struct {
	Method     string
	PublicPath string
	ServiceKey string
	TargetPath string
}

type RouteRegistry struct {
	routes map[string]RouteMapping
}

func NewRouteRegistry() *RouteRegistry {
	return &RouteRegistry{
		routes: make(map[string]RouteMapping),
	}
}

func (r *RouteRegistry) Register(route RouteMapping) {
	key := route.Method + ":" + route.PublicPath
	r.routes[key] = route
}

func (r *RouteRegistry) Resolve(method, path string) (*RouteMapping, bool) {
	key := method + ":" + path

	route, ok := r.routes[key]
	if !ok {
		return nil, false
	}

	return &route, true
}

func (rr *RouteRegistry) RegisterAll() {
	rr.Register(RouteMapping{
		Method:     "POST",
		PublicPath: "/api/auth/user/login",
		ServiceKey: "/auth",
		TargetPath: "/api/v1/user/login",
	})

	rr.Register(RouteMapping{
		Method:     "POST",
		PublicPath: "/api/auth/user/register",
		ServiceKey: "/auth",
		TargetPath: "/api/v1/user/register",
	})
}
