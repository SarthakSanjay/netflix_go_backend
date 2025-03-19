package dto

import "github.com/sarthaksanjay/netflix-go/model"

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MovieSuccessResponse struct {
	Message string       `json:"message"`
	Movie   model.Movies `json:"movie"`
}

type MoviesSuccessResponse struct {
	Message string         `json:"message"`
	Movies  []model.Movies `json:"movies"`
	Total   int            `json:"total"`
}

type ShowSuccessResponse struct {
	Message string     `json:"message"`
	Show    model.Show `json:"show"`
}

type ShowsSuccessResponse struct {
	Message string       `json:"message"`
	Shows   []model.Show `json:"shows"`
	Total   int          `json:"total"`
}
