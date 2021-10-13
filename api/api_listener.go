package api

import (
	"fmt"
	"net/http"

	"github.com/hayrullahcansu/cachy/api/routing"
	"github.com/hayrullahcansu/cachy/data/constants"
	"github.com/hayrullahcansu/cachy/framework/logging"
)

type ApiListener struct {
	mux *http.ServeMux
}

// NewApiListener return new instance of ApiWorker
func NewApiListener() *ApiListener {
	return &ApiListener{
		mux: http.NewServeMux(),
	}
}

func (a *ApiListener) ListenAndServe() {

	a.mux.HandleFunc("/", TestServer)
	// a.mux.HandleFunc("/api/v1/hello", HelloServer)
	a.mux.HandleFunc("/api/v1/cache/", routing.CacheRouteHandler)
	uri := fmt.Sprintf(":%d", constants.ListenPort)
	http.ListenAndServe(uri, a.mux)
}

func TestServer(w http.ResponseWriter, r *http.Request) {
	logging.Infof("Hello, %s!", r.URL.Path[1:])

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
