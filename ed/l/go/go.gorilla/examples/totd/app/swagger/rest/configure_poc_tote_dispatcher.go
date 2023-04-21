// This file is safe to edit. Once it exists it will not be overwritten

package rest

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/to-com/poc-td/app/swagger/rest/operations"
)

//go:generate swagger generate server --target ../../swagger --name PocToteDispatcher --spec ../../../api.yaml --model-package restmodel --server-package rest --principal interface{} --exclude-main

func configureFlags(api *operations.PocToteDispatcherAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PocToteDispatcherAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreateToteAssignmentHandler == nil {
		api.CreateToteAssignmentHandler = operations.CreateToteAssignmentHandlerFunc(func(params operations.CreateToteAssignmentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateToteAssignment has not yet been implemented")
		})
	}
	if api.DeleteToteAssignmentHandler == nil {
		api.DeleteToteAssignmentHandler = operations.DeleteToteAssignmentHandlerFunc(func(params operations.DeleteToteAssignmentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteToteAssignment has not yet been implemented")
		})
	}
	if api.GetConfigsHandler == nil {
		api.GetConfigsHandler = operations.GetConfigsHandlerFunc(func(params operations.GetConfigsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetConfigs has not yet been implemented")
		})
	}
	if api.HealthCheckHandler == nil {
		api.HealthCheckHandler = operations.HealthCheckHandlerFunc(func(params operations.HealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.HealthCheck has not yet been implemented")
		})
	}
	if api.ListToteAssignmentsHandler == nil {
		api.ListToteAssignmentsHandler = operations.ListToteAssignmentsHandlerFunc(func(params operations.ListToteAssignmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.ListToteAssignments has not yet been implemented")
		})
	}
	if api.UpdateConfigHandler == nil {
		api.UpdateConfigHandler = operations.UpdateConfigHandlerFunc(func(params operations.UpdateConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateConfig has not yet been implemented")
		})
	}
	if api.UpdateToteAssignmentHandler == nil {
		api.UpdateToteAssignmentHandler = operations.UpdateToteAssignmentHandlerFunc(func(params operations.UpdateToteAssignmentParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.UpdateToteAssignment has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
