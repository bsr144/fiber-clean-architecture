package middlewares

import (
	"errors"
	"konsulin-service/internal/pkg/constvars"
	"konsulin-service/internal/pkg/exceptions"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	code := constvars.StatusInternalServerError
	clientMessage := constvars.ErrClientSomethingWrongWithApplication

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		clientMessage = fiberErr.Message
		logrus.WithFields(logrus.Fields{
			"location": logrus.Fields{
				"file":          constvars.ErrFileLocationUnknown,
				"line":          constvars.ErrLineLocationUnknown,
				"function_name": constvars.ErrFunctionNameUnknown,
			},
		}).Error(fiberErr.Message)
	}

	var customErr *exceptions.CustomError
	if errors.As(err, &customErr) {
		code = customErr.StatusCode
		clientMessage = customErr.ClientMessage
		logrus.WithFields(logrus.Fields{
			"location": logrus.Fields{
				"file":          customErr.Location.File,
				"line":          customErr.Location.Line,
				"function_name": customErr.Location.FunctionName,
			},
		}).Error(customErr.DevMessage)
	} else {
		logrus.Error(err)
	}

	ctx.Set(constvars.HeaderContentType, constvars.MIMEApplicationJSON)
	return ctx.Status(code).JSON(fiber.Map{
		"status_code": code,
		"success":     false,
		"message":     clientMessage,
	})
}
