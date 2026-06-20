package gateway

type RouteMapping struct {
	Method      string
	PublicPath  string
	ServiceKey  string
	TargetPath  string
	RequireAuth bool // Whether this route requires API key verification
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

func (r *RouteRegistry) RegisterAll() {
	// Public auth routes (no authentication required)
	r.Register(RouteMapping{
		Method:      "POST",
		PublicPath:  "/api/auth/user/login",
		ServiceKey:  "/auth",
		TargetPath:  "/api/v1/user/login",
		RequireAuth: false,
	})

	r.Register(RouteMapping{
		Method:      "POST",
		PublicPath:  "/api/auth/user/register",
		ServiceKey:  "/auth",
		TargetPath:  "/api/v1/user/register",
		RequireAuth: false,
	})

	// Protected auth routes (require API key verification)
	r.Register(RouteMapping{
		Method:      "GET",
		PublicPath:  "/api/auth/api-key/verify",
		ServiceKey:  "/auth",
		TargetPath:  "/api/v1/api-key/verify",
		RequireAuth: true,
	})

	// Protected upload service routes (require API key verification)
	r.Register(RouteMapping{
		Method:      "GET",
		PublicPath:  "/api/upload/ping",
		ServiceKey:  "/upload",
		TargetPath:  "/ping",
		RequireAuth: true,
	})
}
