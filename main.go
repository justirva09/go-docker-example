package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Printf("can't load env: %+v", err)
    }

    dbUser := os.Getenv("DATABASE_USER")
    dbPass := os.Getenv("DATABASE_PASSWORD")
    dbHost := os.Getenv("DATABASE_HOST")
    dbPort := os.Getenv("DATABASE_PORT")
    dbName := os.Getenv("DATABASE_NAME")
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

    log.Printf("dsn: %s", dsn)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("can't connect to mysql %+v", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatalf("can't ping mysql: %+v", err)
    }

    blockChan := make(chan bool)

    fmt.Println("connected to mysql")

    <-blockChan
}
