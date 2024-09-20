package main

import (
	"backend-test/internal/di"
	"backend-test/internal/shutdown"
	"backend-test/pkg/http"
	"backend-test/pkg/logger"
)

func main() {
	err := di.GetContainer().Invoke(func(server *http.Server, shutdown *shutdown.ShutdownManager) {
		go shutdown.EnableSignalHandling()

		server.Start()
	})
	if err != nil {
		logger.GetLogger().Fatal("main", "main", err)
	}

}
