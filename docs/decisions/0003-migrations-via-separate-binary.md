# ADR 0003 — Migrations via separate `cmd/migrate/` binary

- **Datum:** 2026-05-02
- **Status:** Accepted

## Context

Memorie heeft persistent state in SQLite (per ADR-0001) en datamodel-evolutie is verwacht (Person/Place/Event/Relationship + later memories_index, en latere tabellen voor curatie/AI-output). Dat vraagt een migration-strategie. Twee paden overwogen:

- **A. Embedded Goose** — Goose als library in `cmd/memorie/main.go`. App-start runt pending migrations, dan start HTTP-server.
- **B. Separate Goose CLI** — `cmd/migrate/` als tweede binary in dezelfde image. Migrations expliciet door operator gerund vóór app-deploy.

## Beslissing

**Pad B: separate `cmd/migrate/` binary.**

Concreet:
- Migrations in `migrations/` als plain `.sql` files met Goose-syntax (`-- +goose Up` / `-- +goose Down`).
- Nieuwe binary `cmd/migrate/main.go` gebaseerd op `github.com/pressly/goose/v3` + `modernc.org/sqlite` (per ADR-0001).
- Beide binaries in **één** Docker image; default `ENTRYPOINT` blijft `["/memorie"]`. Migrations runnen via `docker exec memorie /migrate up` of via een one-shot met overridden entrypoint.
- App-startup runt **geen** migrations. Bewuste keuze: geen mid-startup-corrupte-state mogelijk; rollback-pad expliciet.
- Schema-change-flow: write migration → run `/migrate up` op productie → deploy app-code → bij issue: write reverse migration of `/migrate down`.
- Lokale dev: `go run ./cmd/migrate up` tegen lokale SQLite-file.

## Personas-jury

| | Oordeel | Reden |
|---|---|---|
| Marc | Wint licht | Explicieter, schema-changes apart auditeerbaar; past bij "weet wat ik draai" |
| Sanne | Neutraal | Onzichtbaar achter Portainer-redeploy; geen UX-verschil |
| Tom | Neutraal | Nog steeds 1 container; 2 binaries in image is geen ergonomie-issue |
| Lisa | Neutraal | Install-flow ongewijzigd |

## Bewust verlies

- **Eén extra step bij schema-deploys.** Migrations zijn niet auto-getriggerd, dus een schema-change vereist `docker exec memorie /migrate up` voordat de nieuwe app-code start. Acceptabel: schema-changes zijn zeldzaam in v1, en alle ops loopt via Claude Code (1 extra bash-call). Latere automatisering kan via een nieuwe ADR.
- **~5 MB extra image-grootte** door tweede binary. Verwaarloosbaar binnen distroless-static doel.
- **Geen ingebouwde rollback bij failed migration** — operator moet handmatig reageren. Acceptabel voor v1 (1 family, niet zaakkritiek).

## Afgewezen alternatieven

**A. Embedded Goose in main binary:**
- Voordeel: simpeler, één binary, geen extra deploy-step.
- Nadelen die de doorslag gaven:
  - App-start faalt bij gebroken migration → mogelijk lege/corrupte DB als boot mid-run crasht.
  - Geen rollback zonder app-restart of binary-rebuild.
  - Schema-evolutie is een ops-handeling; impliciet maken voelt als verlies van controle.
- Conflicteert mild met "MVP-first, blokkades-voor-toekomst meewegen": pad B houdt expliciet pad open voor latere CI-automatisering (bv. pre-deploy migration-step in GH Actions) zonder dat de app-binary verandert.

## Consequenties

- Nieuwe directory: `cmd/migrate/main.go`.
- Dockerfile bouwt beide binaries; default-entrypoint blijft `/memorie`.
- Eerste migration (Person/Place/Event/Relationship per Linear GJA-61) wordt `migrations/00001_init_schema.sql`.
- Goose-driver = `sqlite` (modernc-compatible).
- Documentatie: README v0 krijgt een korte sectie "Schema migrations" met de drie commando's (lokaal, prod-exec, prod-one-shot).
- Toekomst: pre-deploy migration-step in CI = aparte ADR zodra schema-cadens dat rechtvaardigt.
