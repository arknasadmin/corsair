package middleware

  import (
  	"net/http"
  	"strings"

  	"github.com/bastienwirtz/corsair/config"
  )

  // TrailingSlash middleware only adds trailing slash for exact endpoint path matches.
  // This prevents Go's http.ServeMux from issuing 301 redirects while not breaking
  // sub-paths that don't want trailing slashes.
  func TrailingSlash(endpoints []config.Endpoint) func(http.Handler) http.Handler {
  	return func(next http.Handler) http.Handler {
  		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  			path := r.URL.Path

  			// Only add trailing slash if path exactly matches an endpoint
  			for _, ep := range endpoints {
  				epPath := strings.TrimSuffix(ep.Path, "/")
  				if path == epPath {
  					newURL := *r.URL
  					newURL.Path = path + "/"
  					r.URL = &newURL
  					break
  				}
  			}

  			next.ServeHTTP(w, r)
  		})
  	}
  }

