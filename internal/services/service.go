package services

import (
	"github.com/fratschi/go-microservice/internal"
	"github.com/rs/zerolog/log"
)

type Data struct {
	ID   string   `json:"id"`
	Data []string `json:"data"`
}

type Service interface {
	Do(id string)
	DoPost(data *Data) error
}

func NewService(config *internal.Config) Service {
	return &service{
		config: config,
	}
}

type service struct {
	config *internal.Config
}

func (s *service) Do(id string) {
	log.Info().Str("id", id).Msg("Service is doing something")
}

func (s *service) DoPost(data *Data) error {
	log.Info().Str("id", data.ID).Strs("data", data.Data).Msg("Service is doing something")
	return nil
}
