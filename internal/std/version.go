package std

import (
	"github.com/go-chi/chi"
	"io"
	"net/http"
)

// can be set from outside within go build command
// -ldflags="-X github.com/fratschi/go-microservice/internal/std.Version=${SERVICE_VERSION}"
var Version string

func HandleVersionV1(router chi.Router) {
	router.MethodFunc("GET", "/v1/version", func(writer http.ResponseWriter, request *http.Request) {
		requestedContentType := request.Header.Get("Content-Type")
		if requestedContentType == "text/plain" {
			writer.Header().Set("Content-Type", "text/plain")
			io.WriteString(writer, Version)
		} else {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			io.WriteString(writer, "{\"version\": \""+Version+"\"}\n")
		}
	})
}
