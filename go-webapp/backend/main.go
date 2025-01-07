package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "go-webapp/backend/config"
    "go-webapp/backend/routes"
)

func main() {
    // Konfiguration laden
    config.LoadConfig()

    // Router initialisieren
    r := mux.NewRouter()

    // Routen konfigurieren
    routes.RegisterRoutes(r)

    // Server starten
    log.Println("Server läuft auf Port:", config.AppConfig.Port)
    if err := http.ListenAndServe(":"+config.AppConfig.Port, r); err != nil {
        log.Fatal("Fehler beim Starten des Servers:", err)
    }
}