package router

type Strategy string

type RouterIterator interface {
	Next() *RouterConfig
}
