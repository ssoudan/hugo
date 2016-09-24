package logging

import (
	"github.com/op/go-logging"

	"os"
)

// Log creates a new named Logger
func Log(name string) *logging.Logger {
	return logging.MustGetLogger(name)
}

func init() {
	colorFormat := logging.MustStringFormatter(
		"%{color:bold}%{time:15:04:05.000} %{level:-6s} [%{module}] %{shortfunc:.10s} ▶ %{id:03x}%{color:reset} %{message}",
	)

	backend := logging.NewLogBackend(os.Stderr, "", 0)

	backendFormatter := logging.NewBackendFormatter(backend, colorFormat)

	// Only errors and more severe messages should be sent to backend1
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.DEBUG, "")

	logging.SetBackend(backendLeveled)
	return
}
