package models

import (
	"encoding/json"
	"strings"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	headers    map[string]string
}

func CreateResponse(statusCode int, data interface{},
	message string) *Response {
	return BuildResponse().
		SetStatusCode(statusCode).
		SetData(data).
		SetMessage(message)
}

func BuildResponse() *Response {
	return &Response{
		headers: map[string]string{"content-type": "application/json", "Access-Control-Allow-Origin": "*"},
	}
}

func (r *Response) SetStatusCode(statusCode int) *Response {
	r.StatusCode = statusCode
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) AddHeader(key, value string) *Response {
	r.headers[strings.ToLower(key)] = value
	return r
}

func (r *Response) RemoveHeader(key string) *Response {
	key = strings.ToLower(key)
	if _, ok := r.headers[key]; ok {
		delete(r.headers, key)
	}

	return r
}

func (r *Response) RemoveHeaders() *Response {
	r.headers = map[string]string{}

	return r
}

func (r *Response) GetHeader(key string) string {
	return r.headers[strings.ToLower(key)]
}

func (r *Response) GetHeaders() map[string]string {
	return r.headers
}

func (r *Response) ToBytes() []byte {
	bytes, _ := json.Marshal(r)
	return bytes
}
