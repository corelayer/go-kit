/*
 * Copyright 2024 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func NewHttpServer(address string, port int, handler http.Handler) *HttpServer {
	return &HttpServer{
		Server: &http.Server{
			Addr:    address + ":" + strconv.Itoa(port),
			Handler: handler},
	}
}

type HttpServer struct {
	*http.Server
}

func (s *HttpServer) RunServer(ctx context.Context) {
	// HttpServer run context
	serverCtx, serverStopCtx := context.WithCancel(ctx)

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Listen for signal notification and run start listening
	slog.Debug("server starting", "address", s.Addr)
	go s.shutdown(&serverCtx, serverStopCtx, &sig)
	go s.listenAndServe()

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

func (s *HttpServer) shutdown(c *context.Context, cancelFunc context.CancelFunc, sig *chan os.Signal) {
	<-*sig

	// Shutdown signal with grace period of 30 seconds
	shutdownCtx, cancel := context.WithTimeout(*c, 30*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			slog.Error("graceful shutdown timed out", "address", s.Addr, "error", ctx.Err())
			os.Exit(1)
		}
	}(shutdownCtx)

	slog.Debug("shutting down server", "address", s.Addr)
	// Trigger graceful shutdown
	err := s.Server.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error("could not shutdown server", "address", s.Addr, "error", err)
	}

	// Call parent context cancel function to complete graceful exit
	cancelFunc()
}

func (s *HttpServer) listenAndServe() {
	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("could not start server", "address", s.Addr, "error", err)
	}
}
