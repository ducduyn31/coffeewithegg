package shutdown

import (
	"context"
	"os/signal"
	"syscall"
)

var (
	signalCtx    context.Context
	signalCancel context.CancelFunc
)

func InitGracefulShutdown() {
	signalCtx, signalCancel = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
}

func MainContext() context.Context {
	return signalCtx
}

func ForceShutdown() {
	signalCancel()
}
