package http

import "problem1/models"

type Response struct {
	Content string          `json:"content"`
	Result  []models.Friend `json:"result"`
	Total   int             `json:"total"`
}

func NewResponse(content string, result []models.Friend, total int) *Response {
	return &Response{content, result, total}
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

type ErrorResponse struct {
	Error ErrorData `json:"error"`
}

func NewErrorResponse(id, code, title, detail, info string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorData{
			Parameter: Parameter{Id: id},
			Code:      code,
			Title:     title,
			Detail:    detail,
			Info:      info,
		},
	}
}
