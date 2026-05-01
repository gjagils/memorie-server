# ADR 0001 — Stack: Go + SQLite + single binary

- **Datum:** 2026-05-01
- **Status:** Accepted

## Context

Memorie draait op Synology-NAS-en met beperkte resources (vaak 2 GB RAM, ARM of x86), naast bestaande containers (Immich, Plex, Home Assistant). Doel uit roadmap: 5-minuten-installatie. Mogelijke distributie-vorm in toekomst: Synology Package Center (SPK).

De gebruiker beheert en ontwikkelt Memorie 100% via Claude Code (geen lokale DBeaver/TablePlus, geen psql-prompt). DB-inspectie gaat via `sqlite3` CLI in Bash op de Synology.

In een eerdere sessie was kort overwogen: Python/FastAPI + Postgres. Bij het uitwerken bleek dat dat botst met de install-belofte (multi-container), de geheugen-footprint, multi-arch builds (C-extensions) en het single-binary-pad voor SPK.

## Beslissing

- **Taal:** Go — single statisch gecompileerde binary
- **DB:** SQLite met WAL-mode; Litestream voor continue replicatie (later, niet in v0)
- **SQLite driver:** voorkeur `modernc.org/sqlite` (CGO-loos, makkelijk multi-arch). Pas overstappen naar `mattn/go-sqlite3` als een feature-gap dwingt.
- **Migrations:** Goose
- **Image:** distroless of `scratch` base, multi-arch (`linux/amd64`, `linux/arm64`)
- **Container-aantal:** 1 hoofd-container (Memorie zelf). Optionele containers (sqlite-web, Litestream-sidecar) zijn opt-in en niet verplicht voor de basis-installatie.

## Personas-jury

| | Oordeel | Reden |
|---|---|---|
| Marc | Wint | Geen extra Postgres-tuning, 100% air-gapped haalbaar, single container past bij minimal-services-discipline |
| Sanne | Wint | 1 stack-entry in Portainer, 1 env-set, geen 2e troubleshoot-pad |
| Tom | Wint | <30 MB image, snelle pull, 1 container, geen orchestratie nodig |
| Lisa | Wint | "1 ding installeren" past bij haar pijngrens (>1 weekend = afhaker) |

## Bewust verlies

- **Geen Postgres-familiariteit voor de developer.** Acceptabel: ontwikkeling loopt 100% via Claude Code; lokale DB-GUI is irrelevant. SQLite-inspectie via `sqlite3 memorie.db` in Bash.
- **Geen ingebouwde concurrent-multi-writer support.** Acceptabel: 1 family per installatie, write-rate is laag. SQLite handelt makkelijk 10k+ writes/sec via WAL.
- **Schema-migraties iets restrictiever dan Postgres.** Acceptabel: SQLite ≥3.35 (2021) ondersteunt `DROP COLUMN` en `RENAME COLUMN`. Goose wraps het.

## Afgewezen alternatieven

**Python/FastAPI + Postgres:** breekt single-binary, vergroot image 8-10×, +100 MB RAM voor Postgres baseline, multi-arch wheels-gedoe, geen SPK-pad. Ergonomie-voordeel (Postgres-familiariteit, lokale GUI) valt weg omdat ontwikkeling via Claude Code loopt.

**Go + Postgres-sidecar:** middenpad. Verliest "single-binary, 5-min-install, SPK-haalbaar" zonder een persona te winnen die niet al bij Go+SQLite wint. Niet nodig.

## Consequenties

- Alle backend-code in Go.
- DB-toegang via `database/sql` + `modernc.org/sqlite`.
- Geen externe DB-service, geen connection-pool-tuning.
- Backup-strategie wordt later ingericht: Litestream → B2 of Tailscale-peer (latere ADR).
- SPK-package blijft een optioneel distributie-pad voor latere fase.
