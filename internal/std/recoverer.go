package std

import (
	"github.com/rs/zerolog/log"
	"net/http"
	"runtime/debug"
)

// JsonLogRecoverer logs the panic (and a backtrace) as json and returns an HTTP 500 (Internal Server Error) status if possible.
func JsonLogRecoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				log.Error().Bytes("stack", debug.Stack()).Msg("panic")
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
