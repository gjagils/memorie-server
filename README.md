# memorie-server

Privacy-first, self-hosted "on this day" memories-app voor families. Laag bovenop [Immich](https://immich.app/). Single Go binary, SQLite, draait in 1 Docker container.

> ⚠️ Vereist een werkende Immich-installatie ([ADR-0002](docs/decisions/0002-immich-integratie-pure-api-laag.md)). Memorie zelf maakt geen tweede kopie van je foto-archief.

## Status

In ontwikkeling — pre-v0. Backlog leeft in Linear (`Memorie` project, GJA-XX issues).

## Stack

Go single binary · SQLite (WAL) · Goose migrations · Docker multi-arch (amd64+arm64) · AGPL-3.0.
Zie [ADR-0001](docs/decisions/0001-stack-go-sqlite-single-binary.md) voor het waarom.

## Lokaal draaien (development)

```bash
go run ./cmd/memorie
# → http://localhost:8090/health
```

## Deploy

GitHub push → GitHub Actions (multi-arch build, GHCR push) → Tailscale → Portainer API → Synology Docker.
Zie [`.github/workflows/deploy.yml`](.github/workflows/deploy.yml).

## Voor ontwikkelaars / Claude Code

Verplichte leesvolgorde voor architectuur-werk:

1. [`PERSONAS.md`](PERSONAS.md) — 4 personas als jury voor design-keuzes
2. [`docs/decisions/`](docs/decisions/) — ADRs (waarom-deze-keuze)
3. [`CLAUDE.md`](CLAUDE.md) — top-level project-context

Geen ADR = geen architectuurkeuze.

## Licentie

[GNU Affero General Public License v3.0](LICENSE). De AGPL-keuze beschermt tegen managed-hosting copycats; eigen self-hosting voor families is uitdrukkelijk gewenst.
