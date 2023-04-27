package main

import (
	"github.com/fratschi/go-microservice/internal"
	"github.com/fratschi/go-microservice/internal/services"
	"github.com/fratschi/go-microservice/internal/std"
	v1 "github.com/fratschi/go-microservice/pkg/port/v1"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type server struct {
	Config  *internal.Config
	Service services.Service
}

func main() {

	config := internal.LoadConfig()
	config.Validate()
	server := NewServer(&config)

	go func() {
		server.initProbes()
	}()

	go func() {
		server.initRest()
	}()

	std.WaitForTermination()
}

func NewServer(config *internal.Config) *server {
	log.Info().Str("name", config.Service.Name).Str("version", config.Service.Version).Msg("starting services")

	zerolog.SetGlobalLevel(getLogLevel(config))

	service := services.NewService(config)

	return &server{
		Config:  config,
		Service: service,
	}

}

func (s server) initProbes() {

	router := chi.NewRouter()
	probes := std.NewProbeService()
	probes.HandleProbes(router)

	probes.AddLive(func() (string, bool) {
		return "server", true
	})

	log.Info().Int("port", s.Config.Rest.Health).Msg("starting health probes")
	http.ListenAndServe(s.Config.Rest.HealthPortValue(), router)
}

func (s server) initRest() {
	log.Info().Int("port", s.Config.Rest.Port).Msg("starting rest endpoint")

	router := chi.NewRouter()
	router.Use(std.JsonLogRecoverer)

	v1.Handle(router, s.Service, s.Config)
	handler := std.Cors(s.Config.Cors).Handler(router)

	err := http.ListenAndServe(s.Config.Rest.PortValue(), handler)
	if err != nil {
		log.Fatal().Str("error", err.Error()).Msg("Error while starting rest endpoint")
	}
}

func getLogLevel(config *internal.Config) zerolog.Level {
	loglevel := config.Service.LogLevel
	if loglevel == "" {
		log.Info().Msg("No loglevel set, using info")
		return zerolog.InfoLevel
	}

	level, err := zerolog.ParseLevel(loglevel)
	if err != nil {
		log.Err(err).Msg("Cannot parse log level")
		return zerolog.InfoLevel
	}

	return level
}
