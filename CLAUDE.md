# Memorie — Server

Privacy-first, self-hosted "on this day" memories-app voor families. Laag bovenop Immich. iOS-first met SwiftUI client (aparte repo, nog niet aangemaakt).

## Stack (zie ADR-0001)

- Go single static binary
- SQLite + WAL; Litestream-replicatie later
- Goose voor migrations
- Multi-arch Docker (amd64 + arm64), distroless of `scratch` base
- AGPL-3.0

## Verplichte leesvolgorde voor design-/architectuur-werk

Voordat je een architectuur- of UX-keuze maakt of voorstelt:

1. **[`PERSONAS.md`](PERSONAS.md)** — 4 personas (Marc, Sanne, Tom, Lisa) als jury. Per keuze tabel met wint/verliest/neutraal + bewust verlies. De "hoe te gebruiken"-sectie daar is verplichte volgorde, geen suggestie.
2. **[`docs/decisions/`](docs/decisions/)** — Architecture Decision Records. Lees relevante ADRs voor je iets voorstelt; nieuwe keuze of een oude omgooien = nieuwe ADR met persona-jury (zie `docs/decisions/README.md`).

Geen ADR = geen architectuurkeuze. "We doen het wel even zo" is geen geldige werkwijze.

## Workflow

- **Backlog:** Linear-project `Memorie` met issues `GJA-XX`. Slash-command `.claude/commands/backlog-routine.md` pakt issues op met label `claude-ready` + status `Todo`.
- **Repo:** `gjagils/memorie-server` (public, AGPL-3.0).
- **Deploy-doel:** Synology via Portainer + GitHub Actions + Tailscale (zie `~/Github/project-template-synology.md` als referentie; pipeline in deze repo nog in te richten).

## Project-structuur

```
.
├── cmd/memorie/             # main.go entrypoint (alleen /health in v0)
├── internal/
│   ├── photosource/         # PhotoSource interface (ADR-0002) — impl volgt
│   ├── memories/            # memory-card generatie (komt) — Person/Place/Event/Relationship
│   └── http/                # HTTP-handlers (komt)
├── migrations/              # Goose migration files
├── docs/decisions/          # ADRs
├── PERSONAS.md
├── CLAUDE.md
├── Dockerfile               # multi-stage, distroless static, multi-arch
├── docker-compose.yml       # 1 service, named volume voor SQLite
├── .env.example
└── .github/workflows/deploy.yml   # build → GHCR → Tailscale → Portainer
```

## Lokaal draaien

```bash
go run ./cmd/memorie
# → http://localhost:8090/health
```

## Pending (bekend werk)

- iOS-repo nog niet aangemaakt (Linear GJA-57).
- Eerste echte schema-migratie (Person/Place/Event/Relationship per datamodel) — Linear GJA-61.
- `ImmichPhotoSource` implementatie + auth-koppeling (latere ADR voor auth).
- Litestream-replicatie (latere ADR).
