package main

import (
  "database/sql"
  "fmt"
  "log"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "admin"
  password = "ad"
  dbname   = "gachi_akuta"
)

func connectDB() (*sql.DB, error) {

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    return nil, err
  }

  fmt.Println("Conexion exitosa")
  return db, nil
  }
  
  func main() {

    db, err := connectDB()
    if err != nil {
      log.Fatal("Error al conectar a la base de datos: ", err)
    }

    defer db.Close()

  }



