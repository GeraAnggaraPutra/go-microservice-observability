package app

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	zlog "github.com/rs/zerolog/log"
)

func StartServer() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(HttpLoggerMiddleware())
	e.Use(echoprometheus.NewMiddleware(os.Getenv("GRAFANA_SERVICE")))
	e.Use(TrackCustomMetrics)

	e.GET("/metrics", echoprometheus.NewHandler())
	e.GET("/health", healthHandler)
	e.GET("/users", usersHandler)
	e.GET("/error", errorHandler)
	e.GET("/random", randomHandler)

	if err := e.Start(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(c *echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Up and running!",
	})
}

func usersHandler(c *echo.Context) (err error) {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello from /users",
	})
}

func errorHandler(c *echo.Context) (err error) {
	err = errors.New("manual trigger error for testing")

	zlog.Error().Err(err).Msg("An error occurred in /error endpoint")

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "Internal Server Error",
	})
}

func randomHandler(c *echo.Context) (err error) {
	start := time.Now()

	randomSleep := rand.Intn(10) + 1
	time.Sleep(time.Duration(randomSleep) * time.Second)

	statusCodes := []int{200, 201, 400, 404, 500, 503}
	randomStatus := statusCodes[rand.Intn(len(statusCodes))]

	duration := time.Since(start).Seconds()
	requestDuration.WithLabelValues(fmt.Sprintf("%d", randomStatus)).Observe(duration)
	customCounter.WithLabelValues(c.Path(), fmt.Sprintf("%d", randomStatus)).Inc()

	if randomStatus >= 400 {
		err = errors.New("randomly generated error for testing")
		zlog.Error().Err(err).
			Int("status", randomStatus).
			Float64("duration_sec", duration).
			Msg("API error triggered by random generator")
	}

	return c.JSON(randomStatus, map[string]interface{}{
		"delay":  fmt.Sprintf("%d seconds", randomSleep),
		"status": randomStatus,
	})
}
