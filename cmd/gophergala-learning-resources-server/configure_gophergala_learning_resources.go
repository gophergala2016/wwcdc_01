package main

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/gophergala2016/wwcdc_01/db"
	"github.com/gophergala2016/wwcdc_01/models"
	"github.com/gophergala2016/wwcdc_01/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func addLearningResource(lr *models.LearningResource) error {
	if lr == nil {
		return errors.New(500, "lr must be present")
	}
	return db.CreateLearningResource(lr)
}

func configureAPI(api *operations.GophergalaLearningResourcesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	api.AddLearningResourceHandler = operations.AddLearningResourceHandlerFunc(func(params operations.AddLearningResourceParams) middleware.Responder {
		if err := addLearningResource(params.LearningResource); err != nil {
			return operations.NewAddLearningResourceDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
		return operations.NewAddLearningResourceOK().WithPayload(params.LearningResource)
	})
	api.AddReviewHandler = operations.AddReviewHandlerFunc(func(params operations.AddReviewParams) middleware.Responder {
		err := db.AddReview(params.Review)
		if err != nil {
			return operations.NewAddReviewDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
		return operations.NewAddReviewOK().WithPayload(params.Review)
	})
	api.AddScreenshotHandler = operations.AddScreenshotHandlerFunc(func(params operations.AddScreenshotParams) middleware.Responder {
		err := db.AddScreenshot(&params)
		if err != nil {
			return operations.NewAddScreenshotDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
		return operations.NewAddScreenshotOK()
	})
	api.AddUserHandler = operations.AddUserHandlerFunc(func(params operations.AddUserParams) middleware.Responder {
		err := db.AddUser(params.User)
		if err != nil {
			return operations.NewAddUserDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
		return operations.NewAddUserOK()
	})
	api.AuthUserHandler = operations.AuthUserHandlerFunc(func(params operations.AuthUserParams) middleware.Responder {
		authed := db.AuthUser(params.User.Email, params.User.Password)
		if authed {
			return operations.NewAuthUserOK()
		}
		return operations.NewAuthUserDefault(403)
	})
	api.DeleteLearningResourceHandler = operations.DeleteLearningResourceHandlerFunc(func(params operations.DeleteLearningResourceParams) middleware.Responder {
		err := db.DeleteLearningResource(params.ID)
		if err != nil {
			return operations.NewDeleteLearningResourceDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
		return operations.NewDeleteLearningResourceNoContent()
	})
	api.DeleteReviewHandler = operations.DeleteReviewHandlerFunc(func(params operations.DeleteReviewParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteReview has not yet been implemented")
	})
	api.DeleteUserHandler = operations.DeleteUserHandlerFunc(func(params operations.DeleteUserParams) middleware.Responder {
		return middleware.NotImplemented("operation .DeleteUser has not yet been implemented")
	})
	api.FindLanguagesHandler = operations.FindLanguagesHandlerFunc(func() middleware.Responder {
		langs, err := db.FindLanguages()
		if err != nil {
			operations.NewFindLanguagesDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
      ls := models.Languages{langs}
		return operations.NewFindLanguagesOK().WithPayload(&ls)
	})
	api.FindLearningResourceByIDHandler = operations.FindLearningResourceByIDHandlerFunc(func(params operations.FindLearningResourceByIDParams) middleware.Responder {
		learningResource, err := db.FindLearningResourceByID(params.ID)
		if err != nil {
			return operations.NewFindLearningResourceByIDDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
		return operations.NewFindLearningResourceByIDOK().WithPayload(learningResource)
	})
	api.FindLearningResourcesHandler = operations.FindLearningResourcesHandlerFunc(func(params operations.FindLearningResourcesParams) middleware.Responder {
		learningResources := []*models.LearningResource{}
		var err error
		if len(params.Types) == 0 {
			learningResources, err = db.FindLearningResources("all")
			if err != nil {
				return operations.NewFindLearningResourcesDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
			}
		}
		for _, t := range params.Types {
			lr, err := db.FindLearningResources(t)
			if err != nil {
				return operations.NewFindLearningResourcesDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
			}
			learningResources = append(learningResources, lr...)
		}
		return operations.NewFindLearningResourcesOK().WithPayload(learningResources)
	})
	api.FindReviewByIDHandler = operations.FindReviewByIDHandlerFunc(func(params operations.FindReviewByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindReviewByID has not yet been implemented")
	})
	api.FindReviewsHandler = operations.FindReviewsHandlerFunc(func() middleware.Responder {
		return middleware.NotImplemented("operation .FindReviews has not yet been implemented")
	})
	api.FindReviewsForLearningResourceHandler = operations.FindReviewsForLearningResourceHandlerFunc(func(params operations.FindReviewsForLearningResourceParams) middleware.Responder {
		reviews, err := db.FindReviewsForLearningResource(params.ID)
		if err != nil {
			return operations.NewFindReviewsForLearningResourceDefault(500).WithPayload(&models.ErrorModel{Code: 500, Message: err.Error()})
		}
		return operations.NewFindReviewsForLearningResourceOK().WithPayload(reviews)
	})
	api.FindUserByIDHandler = operations.FindUserByIDHandlerFunc(func(params operations.FindUserByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation .FindUserByID has not yet been implemented")
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
