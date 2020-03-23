package resolver

var baseURL = "http://localhost:8080"

type RootResolver struct {
	GetAllResolver
	GetByIdResolver
	PostResolver
}
