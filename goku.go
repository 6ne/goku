package goku

import (
	"github.com/clavinjune/goku/models"
	"net/http"
)

func validateStatusCode(statusCode int) int {
	if http.StatusText(statusCode) == "" {
		return 200
	}
	return statusCode
}

func BuildResponse() *models.Response {
	return models.BuildResponse()
}

func CreateOKResponse() *models.Response {
	return CreateResponseWithCode(http.StatusOK)
}

func CreateCreatedResponse() *models.Response {
	return CreateResponseWithCode(http.StatusCreated)
}

func CreateAcceptedResponse() *models.Response {
	return CreateResponseWithCode(http.StatusAccepted)
}

func CreateBadRequestResponse() *models.Response {
	return CreateResponseWithCode(http.StatusBadRequest)
}

func CreateUnauthorizedResponse() *models.Response {
	return CreateResponseWithCode(http.StatusUnauthorized)
}

func CreateForbiddenResponse() *models.Response {
	return CreateResponseWithCode(http.StatusForbidden)
}

func CreateNotFoundResponse() *models.Response {
	return CreateResponseWithCode(http.StatusNotFound)
}

func CreateMethodNotAllowedResponse() *models.Response {
	return CreateResponseWithCode(http.StatusMethodNotAllowed)
}

func CreateInternalServerErrorResponse() *models.Response {
	return CreateResponseWithCode(http.StatusInternalServerError)
}

func CreateResponseWithCode(statusCode int) *models.Response {
	return CreateResponse(statusCode, nil)
}

func CreateResponse(statusCode int, data interface{}) *models.Response {
	statusCode = validateStatusCode(statusCode)
	message := http.StatusText(statusCode)
	return CreateResponseWithMessage(statusCode, data, message)
}

func CreateResponseWithMessage(
	statusCode int, data interface{}, message string) *models.Response {
	return models.CreateResponse(validateStatusCode(statusCode), data, message)
}

func WriteResponse(rw http.ResponseWriter, response *models.Response) {
	for key, val := range response.GetHeaders() {
		rw.Header().Set(key, val)
	}

	rw.WriteHeader(response.StatusCode)
	rw.Write(response.ToBytes())
}
