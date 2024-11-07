package main

import (
    "fmt"
    "log"
    "gachi-akuta-api/config"
)

func main() {
    // Inicializa la configuración general
    config.Init()

    // Conectar a la base de datos
    db, err := config.InitDB()
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }
    defer db.Close()

    fmt.Println("Aplicación iniciada correctamente")

    // Create tables
    config.CreateTables(db)
}

