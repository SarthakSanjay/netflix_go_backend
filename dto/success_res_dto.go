package dto

import "github.com/sarthaksanjay/netflix-go/model"

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MovieSuccessResponse struct {
	Message string         `json:"message"`
	Movies  []model.Movies `json:"message"`
	Total   int            `json:"total"`
}
