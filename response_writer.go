package simplehttp

import "net/http"

type ResponseWriter struct {
	rw         http.ResponseWriter
	statusCode int
	body       []byte
}

func NewResponseWriter(rw http.ResponseWriter) ResponseWriter {
	return ResponseWriter{
		rw:         rw,
		statusCode: 200,
	}
}

func (r ResponseWriter) Header() http.Header {
	return r.rw.Header()
}

func (r ResponseWriter) Write(body []byte) (int, error) {
	r.body = append(r.body, body...)
	return r.rw.Write(body)
}

func (r ResponseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.rw.WriteHeader(statusCode)
}

func (r ResponseWriter) Status() int {
	return r.statusCode
}

func (r ResponseWriter) Body() []byte {
	return r.body
}
