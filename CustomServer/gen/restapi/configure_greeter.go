// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"CustomServer/gen/restapi/operations"
	"CustomServer/gen/restapi/operations/student"
)

//go:generate swagger generate server --target ../../gen --name Greeter --spec ../../swagger/swagger.yml --exclude-main

func configureFlags(api *operations.GreeterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GreeterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TxtProducer = runtime.TextProducer()

	if api.StudentDeleteStudentHandler == nil {
		api.StudentDeleteStudentHandler = student.DeleteStudentHandlerFunc(func(params student.DeleteStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation student.DeleteStudent has not yet been implemented")
		})
	}
	if api.StudentGetStudentHandler == nil {
		api.StudentGetStudentHandler = student.GetStudentHandlerFunc(func(params student.GetStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation student.GetStudent has not yet been implemented")
		})
	}
	if api.StudentGetStudentListHandler == nil {
		api.StudentGetStudentListHandler = student.GetStudentListHandlerFunc(func(params student.GetStudentListParams) middleware.Responder {
			return middleware.NotImplemented("operation student.GetStudentList has not yet been implemented")
		})
	}
	if api.StudentPostStudentHandler == nil {
		api.StudentPostStudentHandler = student.PostStudentHandlerFunc(func(params student.PostStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation student.PostStudent has not yet been implemented")
		})
	}
	if api.StudentPutStudentHandler == nil {
		api.StudentPutStudentHandler = student.PutStudentHandlerFunc(func(params student.PutStudentParams) middleware.Responder {
			return middleware.NotImplemented("operation student.PutStudent has not yet been implemented")
		})
	}

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
