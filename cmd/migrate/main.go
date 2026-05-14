// Memorie schema migration runner (ADR-0003).
//
// Usage:
//
//	migrate up            apply all pending migrations
//	migrate down          roll back the most recent migration
//	migrate status        list applied/pending migrations
//	migrate version       print current schema version
//	migrate reset         roll back all migrations
//
// Reads DATABASE_PATH from the environment (default /data/memorie.db).
// Ships in the same Docker image as the memorie binary; run via
// `docker exec memorie /migrate up`.
package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"

	"github.com/gjagils/memorie-server/migrations"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: migrate <up|down|status|version|reset>")
	}
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(2)
	}
	cmd := flag.Arg(0)

	if err := run(cmd); err != nil {
		log.Fatalf("migrate %s: %v", cmd, err)
	}
}

func run(cmd string) error {
	dbPath := getenv("DATABASE_PATH", "/data/memorie.db")

	db, err := sql.Open("sqlite", dbPath+"?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)")
	if err != nil {
		return fmt.Errorf("open db at %s: %w", dbPath, err)
	}
	defer db.Close()

	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	ctx := context.Background()
	switch cmd {
	case "up":
		return goose.UpContext(ctx, db, ".")
	case "down":
		return goose.DownContext(ctx, db, ".")
	case "status":
		return goose.StatusContext(ctx, db, ".")
	case "version":
		v, err := goose.GetDBVersionContext(ctx, db)
		if err != nil {
			return err
		}
		fmt.Println(v)
		return nil
	case "reset":
		return goose.ResetContext(ctx, db, ".")
	default:
		return errors.New("unknown command (use up, down, status, version, reset)")
	}
}

func getenv(k, fallback string) string {
	if v, ok := os.LookupEnv(k); ok && v != "" {
		return v
	}
	return fallback
}
