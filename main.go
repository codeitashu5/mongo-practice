package main

import (
	"github.com/sirupsen/logrus"
	"mongoPractice/mongoClient"
	"mongoPractice/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
const shutDownTimeOut = 10 * time.Second

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	srv := server.SetupRoutes()

	// start the mongoDB client
	mongoClient.SetUpMongo()
	logrus.Info(">> mongoClient")

	go func() {
		if err := srv.Run(":8080"); err != nil && err != http.ErrServerClosed {
			logrus.Panicf("Failed to run server with error: %+v", err)
		}
	}()

	logrus.Print("Server started at :8080")
	<-done
	logrus.Info("shutting down server")

	if err := mongoClient.ShutdownMongo(); err != nil {
		logrus.WithError(err).Error("failed to close client connection")
	}

	logrus.Info("<< mongoClient")

	if err := srv.Shutdown(shutDownTimeOut); err != nil {
		logrus.WithError(err).Panic("failed to gracefully shutdown server")
	}
}