package controllers

import "problem1/models"

type Response struct {
	Content string          `json:"content"`
	Result  []models.Friend `json:"result"`
	Total   int             `json:"total"`
}

func newResponse(content string, result []models.Friend, total int) *Response {
	return &Response{content, result, total}
}
