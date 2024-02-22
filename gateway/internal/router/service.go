package router

var _ IRouter = &RouterService{}

type RouterService struct{}

func NewRouterService() *RouterService {
	return &RouterService{}
}
