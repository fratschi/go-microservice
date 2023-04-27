package std

import (
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

// WaitForTermination waits for the TERM or INT signal to stop the services.
func WaitForTermination() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	rec := <-sig
	log.Info().Str("signal", rec.String()).Msg("terminating services")
}

// WaitForTerminationCallback waits for the TERM or INT signal to stop the services.
// If signal is received, the callback is executed
func WaitForTerminationCallback(callback func(signal string)) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	rec := <-sig
	callback(rec.String())
}
