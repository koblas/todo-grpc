package bufcutil

import (
	"net/http"
	"net/url"
)

func AddPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = prefix + r.URL.Path
		r2.URL.RawPath = prefix + r.URL.RawPath
		h.ServeHTTP(w, r2)
	})
}

func Rewrite(from, to string, handler http.Handler) http.Handler {
	return http.StripPrefix(from, AddPrefix(to, handler))
}

func RewriteMux(from, to string, handler http.Handler) (string, http.Handler) {
	return from, Rewrite(from, to, handler)
}
