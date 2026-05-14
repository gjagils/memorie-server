// Package migrations embeds the Goose SQL migration files so they ship
// inside the migrate binary (ADR-0003). All schema changes live as plain
// .sql files in this directory.
package migrations

import "embed"

//go:embed *.sql
var FS embed.FS
