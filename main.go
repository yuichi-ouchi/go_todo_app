package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yuichi-ouchi/go_todo_app/config"
	"golang.org/x/sync/errgroup"
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

	// 引数でnet.Litnerを受け取る（addr フィールドは使わない）
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//! For test command Interrupt signal.
			time.Sleep(5 * time.Second)
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)

	//別ゴルーチンでHTTPサーバを起動する
	eg.Go(func() error {
		// ErrServerClosedは http.Server.Shutdownが正常終了したことを示す（異常ではない）
		if err := s.Serve(l); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	//チェネルからの終了通知を待機
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	// Goメソッドで起動した別のゴルーチン終了を待つ
	return eg.Wait()
}
