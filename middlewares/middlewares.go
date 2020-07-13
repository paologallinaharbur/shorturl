package middlewares

import (
	"github.com/paologallinaharbur/shorturl/storage"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strings"
)

//RedirectMiddleware is in charge of redirecting short URLs adding the protocol if needed
func RedirectMiddleware(handler http.Handler, db storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.HasPrefix(r.URL.Path, "/short-") {
			url, err := db.Read(strings.Trim(r.URL.Path, "/"))
			if err != nil {
				http.Redirect(w, r, "http://www.notfound.com", http.StatusSeeOther)
				return
			}

			if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
				url = "https://" + url
			}
			http.Redirect(w, r, url, http.StatusSeeOther)

			numberRedirections.Inc()
			return
		}

		handler.ServeHTTP(w, r)

	})
}

//UIMiddleware is in charge of redirecting to the swagger-ui page
func UIMiddleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Shortcut helpers for swagger-ui
		if r.URL.Path == "/swagger-ui" || r.URL.Path == "/" {
			http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
			return
		}
		// Serving ./swagger-ui/
		if strings.Index(r.URL.Path, "/swagger-ui/") == 0 {
			http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("swagger-ui"))).ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)

	})
}

//PrometheusMiddleware is in charge of redirecting to the prometheus metrics
func PrometheusMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if strings.Index(r.URL.Path, "/metrics") == 0 {
			promhttp.Handler().ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

//Prometheus metrics
func init() {
	prometheus.MustRegister(numberRedirections)
}

var (
	numberRedirections = prometheus.NewCounter(prometheus.CounterOpts{
		Name:        "url_redirected",
		Help:        "Number of redirect.",
		ConstLabels: prometheus.Labels{"version": "1.0.0"},
	})
)
