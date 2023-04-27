package internal

import (
	"github.com/fratschi/go-microservice/internal/std"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Config struct {
	Service Service        `yaml:"services"`
	Rest    Rest           `yaml:"rest"`
	Cors    std.CorsConfig `yaml:"cors"`
}

type Service struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	LogLevel    string `yaml:"logLevel"`
}

type Rest struct {
	Port   int
	Health int
}

func (s Rest) PortValue() string {
	if s.Port == 0 {
		log.Warn().Msg("no port configured, using default port 8080")
		return ":8080"
	}
	return ":" + strconv.Itoa(s.Port)
}

func (s Rest) HealthPortValue() string {
	if s.Port == 0 {
		log.Warn().Msg("no health port configured, using default port 8081")
		return ":8081"
	}
	return ":" + strconv.Itoa(s.Health)
}

func LoadConfig() (c Config) {
	c = Config{}
	err := std.NewConfigLoader().LoadConfig(&c)
	if err != nil {
		log.Err(err).Msg("cannot load config")
	}
	return
}

func LoadConfigData(configData []byte) Config {
	c := Config{}
	err := std.NewConfigLoader().LoadConfigData(configData, &c)
	if err != nil {
		log.Err(err).Msg("cannot load config")
	}
	return c

}

func (c Config) Validate() {
	if c.Service.Name == "" {
		log.Warn().Msg("services name is missing")
	}
	if c.Service.Version == "" {
		log.Warn().Msg("services version is missing")
	}
	if c.Service.Description == "" {
		log.Warn().Msg("services description is missing")
	}

	if c.Rest.Port == 0 {
		log.Warn().Msg("rest port is missing")
	}

	if c.Rest.Health == 0 {
		log.Warn().Msg("rest health port is missing")
	}

	if len(c.Cors.AllowedMethods) == 0 {
		log.Warn().Msg("no allowed methods configured")
	}
}
