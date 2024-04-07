package main

import (
	"context"
	"flag"
	"fmt"
	"helloworldapp/config"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)


const (
    dialect = "pgx"
    dbString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
)

var (
    flags = flag.NewFlagSet("migrate", flag.ExitOnError)
    dir = flags.String("dir", "migrations", "directory with migration files (default is 'migrations')")
)

func main() {
    flags.Usage = usage 
    flags.Parse(os.Args[1:])
    ctx := context.Background()

    args := flags.Args()
    if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
        flags.Usage()
        return
    }

    command := args[0]
    c := config.NewDB()
    dbString := fmt.Sprintf(dbString, c.Host, c.Port, c.Username, c.Password, c.DBName, c.Port)

    db, err := goose.OpenDBWithDriver(dialect, dbString)
    if err != nil {
        log.Fatalf("goose: failed to open DB: %v\n", err)
    }

    defer func() {
        if err := db.Close(); err != nil {
            log.Fatalf("goose: failed to close DB: %v\n", err)
        }
    }()

    if err := goose.RunContext(ctx, command, db, *dir, args[1:]...); err != nil {
        log.Fatalf("goose run: %v", err)
    }
}

func usage() {
    fmt.Print(usagePrefix)
    flags.PrintDefaults()
    fmt.Print(usageCommands)
}

var (
    usagePrefix = `Usage: goose [OPTIONS] COMMAND
Examples:
    migrate status
`
    usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations
`
)
