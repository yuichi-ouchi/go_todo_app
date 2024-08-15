package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, mux http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: mux},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// for Graceful Shutdown
	// if process kill signal or server shutdown signal is received.
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Kill, os.Interrupt)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// ErrServerClosedは
		// http.Server.Shutdownが正常終了したことを示す（異常ではない）
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to close: %+v", err)
	}
	// wait Graceful Shutdown
	return eg.Wait()
}
