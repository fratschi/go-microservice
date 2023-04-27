package v1

import (
	"encoding/json"
	"github.com/fratschi/go-microservice/internal"
	"github.com/fratschi/go-microservice/internal/services"
	"github.com/fratschi/go-microservice/internal/std"
	"github.com/go-chi/chi"
	"net/http"
)

func Handle(router chi.Router, service services.Service, config *internal.Config) {
	// provide version endpoint
	std.HandleVersionV1(router)

	router.MethodFunc("GET", "/v1/test", func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte("{\"test\": \"test\"}"))
		writer.WriteHeader(http.StatusOK)
	})

	router.MethodFunc("GET", "/v1/param/{id}", func(writer http.ResponseWriter, request *http.Request) {

		id := chi.URLParam(request, "id")
		if id == "" {
			std.WriteErrorResponseLog(writer, http.StatusBadRequest, "invalid tenant", nil)
			return
		}

		service.Do(id)

	})

	router.MethodFunc("GET", "/v1/panic", func(writer http.ResponseWriter, request *http.Request) {
		// test for panic recovery middleware
		panic("panic")
	})

	router.MethodFunc("POST", "/v1/post", func(writer http.ResponseWriter, request *http.Request) {

		if request.Header.Get("Content-Type") != "application/json" {
			std.WriteErrorResponseLog(writer, http.StatusBadRequest, "invalid content type", nil)
			return
		}

		req := &PostRequest{}
		err := json.NewDecoder(request.Body).Decode(&req)
		if err != nil {
			std.WriteErrorResponseLog(writer, http.StatusBadRequest, "invalid json", err)
			return
		}

		d := &services.Data{
			ID:   req.ID,
			Data: req.Values,
		}

		err = service.DoPost(d)

		if err != nil {
			std.WriteMessageResponse(writer, http.StatusInternalServerError, "internal error")
		}
	})

}
