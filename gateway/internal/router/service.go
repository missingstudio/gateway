package router

var _ RouterIterator = &RouterService{}

type RouterService struct {
	stack []*RouterConfig
}

func NewRouterService(root *RouterConfig) RouterIterator {
	stack := []*RouterConfig{root}
	return &RouterService{stack}
}

func (rc *RouterService) Iterator() RouterIterator {
	return rc
}

func (rc *RouterService) HasNext() bool {
	return len(rc.stack) > 0
}

func (rc *RouterService) Next() *RouterConfig {
	if !rc.HasNext() {
		return nil
	}

	current := rc.stack[len(rc.stack)-1]
	rc.stack = rc.stack[:len(rc.stack)-1]

	for i := len(current.Targets) - 1; i >= 0; i-- {
		// update stack using strategy
		rc.stack = append(rc.stack, &current.Targets[i])
	}
	return current
}
