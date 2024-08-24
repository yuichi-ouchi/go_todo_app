package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/yuichi-ouchi/go_todo_app/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server:%v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// for Graceful Shutdown
	// if process kill signal or server shutdown signal is received.
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Kill, os.Interrupt)
	defer stop()

	// 環境変数から設定を取得
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// ポート番号取得
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux, cleanup, err := NewMux(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()

	s := NewServer(l, mux)
	return s.Run(ctx)
}
