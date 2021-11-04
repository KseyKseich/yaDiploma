package main

import (
	"context"

	"github.com/AlehaWP/yaDiploma.git/internal/config"
	"github.com/AlehaWP/yaDiploma.git/internal/database"
	"github.com/AlehaWP/yaDiploma.git/internal/server"
	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
	"github.com/AlehaWP/yaDiploma.git/pkg/os_signal"
)

func main() {
	logger.NewLogs()
	defer logger.Close()
	logger.Info("Старт сервера")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config.NewConfig()

	sDB := database.OpenDBConnect()
	defer sDB.Close()

	go os_signal.HandleQuit(cancel)

	s := new(server.Server)
	s.ServerDB = sDB
	s.Start(ctx)

	<-ctx.Done()

}
