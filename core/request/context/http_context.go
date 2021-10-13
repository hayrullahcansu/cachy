package context

import "net/http"

type HttpContext struct {
	W http.ResponseWriter
	R *http.Request
}
