package staticbin

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
)

// Static returns a middleware handler that serves static files in the given directory.
func Static(dir string, Asset func(string) ([]byte, error), options ...Options) martini.Handler {
	if Asset == nil {
		return func() {}
	}

	opt := retrieveOptions(options)

	modtime := time.Now()

	return func(res http.ResponseWriter, req *http.Request, log *log.Logger) {
		if req.Method != "GET" && req.Method != "HEAD" {
			return
		}

		path := req.URL.Path

		b, err := Asset(dir + path)

		if err != nil {
			return
		}

		if !opt.SkipLogging {
			log.Println("[Static] Serving " + path)
		}

		http.ServeContent(res, req, path, modtime, bytes.NewReader(b))
	}
}
