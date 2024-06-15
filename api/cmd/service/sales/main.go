package main

import (
	"context"
	"fmt"
	"github.com/nicholas/store_service/foundation/logger"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {

	var log *logger.Logger

	events := logger.Events{Error: func(ctx context.Context, rec logger.Record) {

		log.Info(ctx, "******* SEND ALERT *******")
	}}

	traceIDFn := func(ctx context.Context) string {
		return ""
	}
	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES-API", traceIDFn, events)

	ctx := context.Background()
	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "msg", err)
		return
	}

}

func run(ctx context.Context, log *logger.Logger) error {

	fmt.Println("code is run fun.")
	//-------------------------------------------------------------------------
	// GOMAXPROCS
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	// -------------------------------------------------------------------------

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	sig := <-shutdown

	log.Info(ctx, "shutdown", "status", "shutdown started", "signal", sig)

	defer log.Info(ctx, "shutdown", "status", "shutdown finished", "signal", sig)
	return nil
}
