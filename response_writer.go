package simplehttp

import "net/http"

type ResponseWriter interface {
	http.ResponseWriter

	Status() int
	Body() []byte
}

type responseWriter struct {
	rw         http.ResponseWriter
	statusCode int
	body       []byte
}

func NewResponseWriter(rw http.ResponseWriter) responseWriter {
	return responseWriter{
		rw:         rw,
		statusCode: 200,
	}
}

func (r responseWriter) Header() http.Header {
	return r.rw.Header()
}

func (r responseWriter) Write(body []byte) (int, error) {
	r.body = append(r.body, body...)
	return r.rw.Write(body)
}

func (r responseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.rw.WriteHeader(statusCode)
}

func (r responseWriter) Status() int {
	return r.statusCode
}

func (r responseWriter) Body() []byte {
	return r.body
}
