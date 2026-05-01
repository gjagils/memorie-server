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

## Project-structuur (target — ontstaat zodra pipeline en eerste code er staan)

```
.
├── cmd/memorie/             # main.go entrypoint
├── internal/
│   ├── photosource/         # PhotoSource interface + ImmichPhotoSource impl (ADR-0002)
│   ├── memories/            # memory-card generatie, Person/Place/Event/Relationship
│   └── http/                # HTTP-handlers
├── migrations/              # Goose migration files
├── docs/decisions/          # ADRs
├── PERSONAS.md
├── CLAUDE.md
├── Dockerfile
├── docker-compose.yml
└── .github/workflows/deploy.yml
```

## Pending (bekend werk)

- LICENSE staat nog op MIT, moet AGPL-3.0 worden (Linear GJA-56).
- Pipeline (Dockerfile, GH Actions, Portainer-deploy) nog niet ingericht.
- iOS-repo nog niet aangemaakt (Linear GJA-57).
