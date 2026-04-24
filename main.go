package main

import (
  "net/http"
  "fmt"
  "log"
  "simple-clothes-store/configserver"
  "simple-clothes-store/internal/handlers"
  "simple-clothes-store/internal/routes"
)

func main() {
  config, err := configserver.LoadConfig()
  if err != nil {
    log.Fatalf("Faild to load configs: %v", err)
  }
  mux := http.NewServeMux()
  handler := handlers.NewHandler()
  routes.SetupRoutes(mux, handler)
  serverAddress := fmt.Sprintf(":%s", config.ServerPort)
  server := &http.Server{
    Addr: serverAddress,
    Handler: mux,
  }
  fmt.Printf("Server Starting On Port %s\n", serverAddress)
  if err := server.ListenAndServe(); err != nil {
    log.Fatalf("Faild to startting server!: %v", err)
  }
}