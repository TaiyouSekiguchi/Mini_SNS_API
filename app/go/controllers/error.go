package controllers

// Error code
const IdErrorCode = "1001"
const NotFoundCode = "2001"

// Error title
const InvalidRequest = "Invalid Request"
const NotFound = "Not Found"

// Error detail
const IdErrorDetail = "The 'id' query parameter value is not valid."
const NotFoundDetail = "Could not find user with parameter id."

// Reference URL
const InfoUrl = "http://example.com/error.html"

type ErrorResponse struct {
	Error ErrorData `json:"error"`
}

type Parameter struct {
	Id string `json:"id"`
}

type ErrorData struct {
	Parameter Parameter `json:"parameters"`
	Code      string    `json:"code"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Info      string    `json:"info"`
}

func newErrorResponse(data ErrorData) *ErrorResponse {
	return &ErrorResponse{data}
}

func newErrorData(id, code, title, detail string) *ErrorData {
	return &ErrorData{Parameter{Id: id}, code, title, detail, InfoUrl}
}
