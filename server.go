package simplehttp

import "net/http"

type Handler struct {
	Funcs []Middleware
}

type Middleware func(rw ResponseWriter, r *http.Request, next func())

func (h Handler) Use(funcs ...Middleware) {
	h.Funcs = append(h.Funcs, funcs...)
}

func (h Handler) ServeHTTP(httpRw http.ResponseWriter, r *http.Request) {
	rw := NewResponseWriter(httpRw)

	composeMiddlewares(rw, r, h.Funcs)()
}

func composeMiddlewares(rw ResponseWriter, r *http.Request, funcs []Middleware) func() {
	if len(funcs) > 0 {
		return func() {
			funcs[0](rw, r, composeMiddlewares(rw, r, funcs[1:]))
		}
	}

	return func() {}
}
