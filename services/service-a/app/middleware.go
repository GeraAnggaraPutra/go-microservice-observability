package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	zlog "github.com/rs/zerolog/log"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func TrackCustomMetrics(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) (err error) {
		if c.Path() == "/metrics" || c.Path() == "/health" {
			return next(c)
		}

		start := time.Now()

		rec := &statusRecorder{
			ResponseWriter: c.Response(),
			status:         http.StatusOK,
		}

		c.SetResponse(rec)

		err = next(c)

		path := c.RouteInfo().Path
		status := rec.status

		customCounter.WithLabelValues(path, fmt.Sprintf("%d", status)).Inc()

		if err != nil || status >= 400 {
			customErrorCounter.WithLabelValues(path, fmt.Sprintf("%d", status)).Inc()
		}

		requestDuration.WithLabelValues(fmt.Sprintf("%d", status)).Observe(time.Since(start).Seconds())

		return
	}
}

func HttpLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:  true,
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			zlog.Info().
				Str("method", v.Method).
				Str("path", c.RouteInfo().Path).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Err(v.Error).
				Msg("http_request")
			return nil
		},
	})
}
