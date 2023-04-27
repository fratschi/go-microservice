package std

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hash/adler32"
	"net/http"
	"strings"
)

// ErrorResponse is a Go structure that encapsulates error information for the rest interfaces.
type ErrorResponse struct {
	// Code is a short, machine-readable identifier that indicates the type of error that occurred.
	Code string `json:"code"`
	// Message provides more human-readable information about what went wrong.
	Message string `json:"message"`
}

// Bytes returns the json information of the error as bytes
func (e ErrorResponse) Bytes() []byte {
	b, _ := json.Marshal(e)
	return b
}

func fromError(err error) ErrorResponse {
	if err == nil {
		return fromMsg("Unknown error")
	}
	return ErrorResponse{
		Code:    code(err.Error()),
		Message: err.Error(),
	}
}

func fromMsg(msg string) ErrorResponse {
	return ErrorResponse{
		Code:    code(msg),
		Message: msg,
	}
}

// code returns an error code for the given message
func code(s string) string {
	return "E" + strings.ToUpper(fmt.Sprintf("%x", ^adler32.Checksum([]byte(s))))
}

func WriteErrorResponseLog(w http.ResponseWriter, status int, msg string, err error) {
	logByStatus(status, msg, err)
	WriteErrorResponse(w, status, err)
}

func WriteErrorResponse(w http.ResponseWriter, status int, err error) {
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(status)
	if err != nil {
		_, _ = w.Write(fromError(err).Bytes())
	}
}

func WriteMessageResponse(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(fromMsg(msg).Bytes())
}

func WriteMessageResponseLog(w http.ResponseWriter, status int, msg string) {
	logByStatus(status, msg, nil)
	WriteMessageResponse(w, status, msg)
}

func logByStatus(status int, msg string, err error) {
	e := eventByStatus(status)
	code := code(msg)
	if err != nil {
		e.Int("status", status).Str("code", code).Err(err).Msg(msg)
	}
	e.Int("status", status).Str("code", code).Msg(msg)
}

func eventByStatus(status int) *zerolog.Event {
	if status >= 500 {
		return log.Error()
	} else if status >= 400 {
		return log.Info()
	} else {
		return log.Debug()
	}
}
