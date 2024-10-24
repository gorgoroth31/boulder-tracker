package repository

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() {
	// Context
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		<-appSignal
		stop()
	}()

	User := os.Getenv("DBUSER")
	Passwd := os.Getenv("DBPASS")
	Net := "tcp"
	Addr := os.Getenv("DBURL")
	DBName := "BoulderTracker"

	// Use parseTime=true to scan MySQL DATETIME type to Go time.Time
	dsn := (User + ":" + Passwd + "@" + Net + "(" + Addr + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	// Open connection
	OpenDbConnection(ctx, db)
}

func OpenDbConnection(ctx context.Context, db *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Unable to connect to database: %v", err)
	}
}
