package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server:%v", err)
	}
}

func run(ctx context.Context, l net.Listener) error {
	// 引数でnet.Litnerを受け取る（addr フィールドは使わない）
	s := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
