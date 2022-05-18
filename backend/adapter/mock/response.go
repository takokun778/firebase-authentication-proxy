package mock

import (
	"net/http"
)

type ResponseWriter struct {
	Assert func(statusCode int)
}

func (r *ResponseWriter) Header() http.Header {
	return http.Header{}
}

func (r *ResponseWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.Assert(statusCode)
}
