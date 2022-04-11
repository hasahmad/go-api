package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func WriteJSON(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func LogError(logger *log.Logger, r *http.Request, err error) {
	logger.WithFields(log.Fields{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	}).Error(err)
}

func ErrorResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	data := Envelope{"error": message}

	err := WriteJSON(w, status, data, nil)
	if err != nil {
		LogError(logger, r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, err error) {
	LogError(logger, r, err)

	message := "the server encountered a problem and could not process your request"
	ErrorResponse(logger, w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	ErrorResponse(logger, w, r, http.StatusNotFound, message)
}

func NotFoundResponseHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		NotFoundResponse(logger, w, r)
	}
}

func MethodNotAllowedResponseHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
		ErrorResponse(logger, w, r, http.StatusMethodNotAllowed, message)
	}
}

func BadRequestResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(logger, w, r, http.StatusBadRequest, err.Error())
}

func InvalidInputResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	ErrorResponse(logger, w, r, http.StatusUnprocessableEntity, "invalid input")
}

func FailedValidationResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, errors map[string]string) {
	ErrorResponse(logger, w, r, http.StatusUnprocessableEntity, errors)
}

func EditConflictResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	ErrorResponse(logger, w, r, http.StatusConflict, message)
}

func RateLimitExceededResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	ErrorResponse(logger, w, r, http.StatusTooManyRequests, message)
}

func InvalidCredentialsResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	ErrorResponse(logger, w, r, http.StatusUnauthorized, message)
}

func InvalidAuthenticationTokenResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	// help inform or remind the client that we expect them to authenticate using a bearer token.
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	ErrorResponse(logger, w, r, http.StatusUnauthorized, message)
}

func AuthenticationRequiredResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	ErrorResponse(logger, w, r, http.StatusUnauthorized, message)
}

func InactiveAccountResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	ErrorResponse(logger, w, r, http.StatusForbidden, message)
}

func NotPermittedResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	ErrorResponse(logger, w, r, http.StatusForbidden, message)
}
