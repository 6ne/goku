package goku

import (
	"encoding/json"
	"github.com/6ne/goku/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidateStatusCode(t *testing.T) {
	testTable := []struct {
		input    int
		expected int
	}{
		{1, 200}, {2, 200}, {3, 200}, {4, 200}, {5, 200}, {6, 200}, {7, 200},
		{8, 200}, {9, 200}, {10, 200}, {11, 200}, {12, 200}, {13, 200}, {14, 200},

		{103, 200}, {306, 200},

		{100, 100}, {101, 101}, {102, 102},

		{200, 200}, {201, 201}, {202, 202}, {203, 203}, {204, 204}, {205, 205},
		{206, 206}, {207, 207}, {208, 208}, {226, 226},

		{300, 300}, {301, 301}, {302, 302}, {303, 303}, {304, 304}, {305, 305},
		{307, 307}, {308, 308},

		{400, 400}, {401, 401}, {402, 402}, {403, 403}, {404, 404}, {405, 405},
		{406, 406}, {407, 407}, {408, 408}, {409, 409}, {410, 410}, {411, 411},
		{412, 412}, {413, 413}, {414, 414}, {415, 415}, {416, 416}, {417, 417},
		{418, 418}, {421, 421}, {422, 422}, {423, 423}, {424, 424}, {425, 425},
		{426, 426}, {428, 428}, {429, 429}, {431, 431}, {451, 451},

		{500, 500}, {501, 501}, {502, 502}, {503, 503}, {504, 504}, {505, 505},
		{506, 506}, {507, 507}, {508, 508}, {510, 510}, {511, 511},
	}

	for _, test := range testTable {
		actual := validateStatusCode(test.input)

		if actual != test.expected {
			t.Errorf("validateStatusCode(%d): expected %d, actual %d",
				test.input, test.expected, actual)
		}
	}
}

func TestBuildResponse(t *testing.T) {
	expected := &models.Response{}
	actual := BuildResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("BuildResponse(): expected %v, actual %v", expected, actual)
	}
}

func TestCreateOKResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 200,
		Message:    "OK",
	}
	actual := CreateOKResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateOKResponse(): expected %v, actual %v", expected, actual)
	}
}

func TestCreateCreatedResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 201,
		Message:    "Created",
	}
	actual := CreateCreatedResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateCreatedResponse(): expected %v, actual %v", expected,
			actual)
	}
}

func TestCreateAcceptedResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 202,
		Message:    "Accepted",
	}
	actual := CreateAcceptedResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateAcceptedResponse(): expected %v, actual %v", expected,
			actual)
	}
}

func TestCreateBadRequestResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 400,
		Message:    "Bad Request",
	}
	actual := CreateBadRequestResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateBadRequestResponse(): expected %v, actual %v", expected,
			actual)
	}
}

func TestCreateUnauthorizedResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 401,
		Message:    "Unauthorized",
	}
	actual := CreateUnauthorizedResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateUnauthorizedResponse(): expected %v, actual %v", expected,
			actual)
	}
}

func TestCreateForbiddenResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 403,
		Message:    "Forbidden",
	}
	actual := CreateForbiddenResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateForbiddenResponse(): expected %v, actual %v", expected,
			actual)
	}
}

func TestCreateNotFoundResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 404,
		Message:    "Not Found",
	}
	actual := CreateNotFoundResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateNotFoundResponse(): expected %v, actual %v", expected,
			actual)
	}
}

func TestCreateMethodNotAllowedResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 405,
		Message:    "Method Not Allowed",
	}
	actual := CreateMethodNotAllowedResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateMethodNotAllowedResponse(): expected %v, actual %v",
			expected, actual)
	}
}

func TestCreateInternalServerErrorResponse(t *testing.T) {
	expected := &models.Response{
		StatusCode: 500,
		Message:    "Internal Server Error",
	}
	actual := CreateInternalServerErrorResponse()

	statusCodeCase := expected.StatusCode == actual.StatusCode
	dataCase := expected.Data == actual.Data
	messageCase := expected.Message == actual.Message

	if !statusCodeCase || !dataCase || !messageCase {
		t.Errorf("CreateInternalServerErrorResponse(): expected %v, actual %v",
			expected, actual)
	}
}

func TestCreateResponseWithCode(t *testing.T) {
	testTable := []struct {
		statusCode int
		expected   *models.Response
	}{
		{20, &models.Response{
			StatusCode: 200,
			Message:    "OK",
		}},
		{103, &models.Response{
			StatusCode: 200,
			Message:    "OK",
		}},
		{202, &models.Response{
			StatusCode: 202,
			Message:    "Accepted",
		}},
		{400, &models.Response{
			StatusCode: 400,
			Message:    "Bad Request",
		}},
		{401, &models.Response{
			StatusCode: 401,
			Message:    "Unauthorized",
		}},
		{500, &models.Response{
			StatusCode: 500,
			Message:    "Internal Server Error",
		}},
	}

	for _, test := range testTable {
		actual := CreateResponseWithCode(test.statusCode)

		statusCodeCase := actual.StatusCode == test.expected.StatusCode
		messageCase := actual.Message == test.expected.Message

		if !statusCodeCase || !messageCase {
			t.Errorf("CreateResponseWithCode(%d): expected %v, actual %v",
				test.statusCode, test.expected, actual)
		}
	}
}

func TestCreateResponse(t *testing.T) {
	testTable := []struct {
		statusCode int
		expected   *models.Response
	}{
		{20, &models.Response{
			StatusCode: 200,
			Message:    "OK",
		}},
		{103, &models.Response{
			StatusCode: 200,
			Message:    "OK",
		}},
		{202, &models.Response{
			StatusCode: 202,
			Message:    "Accepted",
		}},
		{400, &models.Response{
			StatusCode: 400,
			Message:    "Bad Request",
		}},
		{401, &models.Response{
			StatusCode: 401,
			Message:    "Unauthorized",
		}},
		{500, &models.Response{
			StatusCode: 500,
			Message:    "Internal Server Error",
		}},
	}

	for _, test := range testTable {
		actual := CreateResponse(test.statusCode, nil)

		statusCodeCase := actual.StatusCode == test.expected.StatusCode
		messageCase := actual.Message == test.expected.Message

		if !statusCodeCase || !messageCase {
			t.Errorf("CreateResponse(%d, nil): expected %v, actual %v",
				test.statusCode, test.expected, actual)
		}
	}
}

func TestCreateResponseWithMessage(t *testing.T) {
	testTable := []struct {
		statusCode int
		message    string
		expected   *models.Response
	}{
		{20, "\"lorem", &models.Response{
			StatusCode: 200,
			Message:    "\"lorem",
		}},
		{104, "ipsum", &models.Response{
			StatusCode: 200,
			Message:    "ipsum",
		}},
		{201, "dolor", &models.Response{
			StatusCode: 201,
			Message:    "dolor",
		}},
		{402, "sit", &models.Response{
			StatusCode: 402,
			Message:    "sit",
		}},
		{404, "amet", &models.Response{
			StatusCode: 404,
			Message:    "amet",
		}},
		{502, "\"", &models.Response{
			StatusCode: 502,
			Message:    "\"",
		}},
	}

	for _, test := range testTable {
		actual := CreateResponseWithMessage(test.statusCode, nil,
			test.message)

		statusCodeCase := actual.StatusCode == test.expected.StatusCode
		messageCase := actual.Message == test.expected.Message

		if !statusCodeCase || !messageCase {
			t.Errorf(
				"CreateResponseWithMessage(%d, nil, %s): expected %v, actual %v",
				test.statusCode, test.message, test.expected, actual)
		}
	}
}

func TestWriteResponse(t *testing.T) {
	request, err := http.NewRequest("GET", "foobar", nil)

	if err != nil {
		t.Fatal(err)
	}

	expected := CreateResponseWithCode(200)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		WriteResponse(rw, expected)
	})

	handler.ServeHTTP(recorder, request)

	actual := &models.Response{}

	buffer, _ := ioutil.ReadAll(recorder.Body)

	json.Unmarshal(buffer, actual)

	if expected.StatusCode != actual.StatusCode {
		t.Errorf("WriteResponse.Code: expected %v, actual %v", actual.StatusCode,
			expected.StatusCode)
	}

	if expected.Data != actual.Data {
		t.Errorf("WriteResponse.Data: expected %v, actual %v", actual.Data,
			expected.Data)
	}

	if expected.Message != actual.Message {
		t.Errorf("WriteResponse.Message: expected %v, actual %v", actual.Message,
			expected.Message)
	}

	for key, val := range recorder.HeaderMap {
		if expected.GetHeader(key) != val[0] {
			t.Errorf("WriteResponse.headers: expected %v, actual %v", expected.
				GetHeader(key), val[0])
		}
	}
}
