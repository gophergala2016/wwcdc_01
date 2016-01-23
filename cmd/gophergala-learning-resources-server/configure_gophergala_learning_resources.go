package main

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/freeeve/gophergala/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureAPI(api *operations.GophergalaLearningResourcesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	api.AddLearningResourceHandler = operations.AddLearningResourceHandlerFunc(func(params operations.AddLearningResourceParams) middleware.Responder {
		return middleware.NotImplemented("operation .AddLearningResource has not yet been implemented")
	})
	api.DeleteLearningResourceHandler = operations.DeleteLearningResourceHandlerFunc(func(params operations.DeleteLearningResourceParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteLearningResource has not yet been implemented")
	})
	api.FindLearningResourceByIDHandler = operations.FindLearningResourceByIDHandlerFunc(func(params operations.FindLearningResourceByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindLearningResourceByID has not yet been implemented")
	})
	api.FindLearningResourcesHandler = operations.FindLearningResourcesHandlerFunc(func(params operations.FindLearningResourcesParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindLearningResources has not yet been implemented")
	})

	api.ServerShutdown = func() {}
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
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
