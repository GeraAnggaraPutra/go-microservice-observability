package app

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339

	zlog.Logger = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Str("service", os.Getenv("GRAFANA_SERVICE")).
		Logger()
}
