package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log" // Import the log package
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/RecepieApp/server/app"
	"github.com/RecepieApp/server/runtime"
)

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	flag.Parse()

		// mysql 接続
		dsn, ok := os.LookupEnv("DSN")
		dieFalse(ok, "env DSN not found")
	
		db, err := sql.Open("mysql", dsn)
		dieIf(err)
	
		err = db.Ping()
		dieIf(err)
	
		fmt.Println("connected")

		
	a := app.Application{}

	if err := a.LoadConfigurations(); err != nil {
        log.Fatalf("Failed to load configurations: %v", err)
    }

    if err := runtime.Start(&a); err != nil {
        log.Fatalf("Failed to start the application: %v", err)
    }
}


func dieFalse(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
