// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/paologallinaharbur/shorturl/middlewares"
	"github.com/paologallinaharbur/shorturl/restapi/operations/url"
	"github.com/paologallinaharbur/shorturl/restapi/operations/url_implementation"
	"github.com/paologallinaharbur/shorturl/storage"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/paologallinaharbur/shorturl/restapi/operations"
)

//go:generate swagger generate server --target ../../shorturl --name URLShortener --spec ../swagger-ui/swagger.yml

func configureAPI(api *operations.URLShortenerAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()
	storageDB := storage.NewStorageDB("my.db")

	api.PreServerShutdown = func() { storageDB.Close() }
	api.ServerShutdown = func() {}

	//This weird syntax is in place for two reason, it comes directly from go-swagger and it is useful to
	//inject in a function signature an extra parameter (storageDB)
	api.URLCreateURLHandler = url.CreateURLHandlerFunc(func(params url.CreateURLParams) middleware.Responder {
		return url_implementation.CreateURLHandler(params, storageDB)
	})
	api.URLGetURLHandler = url.GetURLHandlerFunc(func(params url.GetURLParams) middleware.Responder {
		return url_implementation.GetURLHandler(params, storageDB)
	})
	api.URLDeleteURLHandler = url.DeleteURLHandlerFunc(func(params url.DeleteURLParams) middleware.Responder {
		return url_implementation.DeleteURLHandler(params, storageDB)
	})

	h1 := api.Serve(nil)
	h2 := middlewares.PrometheusMiddleware(h1)
	h3 := middlewares.UIMiddleware(h2)
	h4 := middlewares.RedirectMiddleware(h3, storageDB)

	return h4
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func configureFlags(api *operations.URLShortenerAPI) {
}
