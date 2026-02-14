package main

import (
	"context"
	"fmt"
	"github/pm-8/students-api/internal/config"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main(){
	//load config
	cfg := config.MustLoad()
	//database setup
	//router setup
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Welcome to Student's API"))
	})
	//server setup
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	slog.Info("server started", slog.String("address", cfg.Addr))
	done := make(chan os.Signal, 1)
	signal.Notify(done,os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func(){
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start the server:", err)
		}
	}()
	<-done
	slog.Info("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to gracefully shutdown the server:", err)
	}
	slog.Info("server shutdown successfully")
	fmt.Printf("Server Started %s",cfg.Addr)
	


}