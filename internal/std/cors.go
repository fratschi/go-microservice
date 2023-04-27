package std

import (
	"github.com/rs/cors"
)

type CorsConfig struct {
	AllowedOrigins []string `yaml:"allowedOrigins"`
	AllowedHeaders []string `yaml:"allowedHeaders"`
	AllowedMethods []string `yaml:"allowedMethods"`
	ExposedHeaders []string `yaml:"exposedHeaders"`
	MaxAge         int      `yaml:"maxAge"`
}

func Cors(c CorsConfig) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: c.AllowedOrigins,
		AllowedMethods: c.AllowedMethods,
		AllowedHeaders: c.AllowedHeaders,
		ExposedHeaders: c.ExposedHeaders,
		MaxAge:         c.MaxAge,
	})
}
