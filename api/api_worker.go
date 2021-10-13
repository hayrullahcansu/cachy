package api

import (
	"fmt"
	"net/http"

	"github.com/hayrullahcansu/cachy/api/routing"
	"github.com/hayrullahcansu/cachy/data/constants"
	"github.com/hayrullahcansu/cachy/framework/logging"
)

type ApiWorker struct {
	mux *http.ServeMux
}

// NewApiWorker return new instance of ApiWorker
func NewApiWorker() *ApiWorker {
	return &ApiWorker{
		mux: http.NewServeMux(),
	}
}

func (a *ApiWorker) ListenAndServce() {

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
