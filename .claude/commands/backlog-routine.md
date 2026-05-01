---
description: Pak Linear issues op met label claude-ready en werk ze één voor één af
---

# Backlog routine voor Memorie

Sequentieel issues uit Linear oppakken, een PR per issue maken, en aan het eind opruimen. Geschreven om dagelijks of weekend-only te draaien zonder supervisie.

## Stappen

1. **Tijd loggen**: noteer starttijd in CEST.

2. **Issues ophalen**: query Linear via MCP voor:
   - Project: `Memorie`
   - Label: `claude-ready`
   - Status: `Todo`
   - Sorteer op priority desc, dan estimate asc (kleine quick wins eerst).

3. **Per issue**:
   - Set status → `In Progress` in Linear.
   - Maak feature branch: gebruik `gitBranchName` uit Linear (bijv. `gerdjanvangils/gja-58-...`).
   - Lees het issue grondig — **acceptance criteria zijn leidend**.
   - Implementeer + tests.
   - Run `make test` en `make lint` lokaal (of `go test ./...` + `golangci-lint run` als de Makefile er nog niet is).
   - Commit met message: `<type>: <korte beschrijving> (GJA-<id>)`.
   - Push branch + open PR met:
     - Titel = issue titel
     - Body = samenvatting + `Closes GJA-<id>`
   - Set issue status → `In Review`.
   - Voeg PR-link als attachment (Link) toe aan het Linear issue.

4. **Stop-criteria** (in deze volgorde):
   - Alle claude-ready issues afgewerkt → ga naar stap 5.
   - Een issue blokkeert (onduidelijk, blokkers niet opgelost): zet terug op `Todo`, voeg comment "blocked: <reden>", ga naar volgende.
   - 3 opeenvolgende test-failures op één issue: idem — comment + door naar volgende.

5. **Opruimen** (alleen als geen open Todo-issues meer met label `claude-ready`):
   - `git fetch -p`
   - Lijst alle remote branches: `git branch -r`
   - Voor elke branch zonder open PR en zonder commits ahead van `main`:
     `git push origin --delete <branch>`
   - Lokale branches die remote weg zijn worden door `git fetch -p` al opgeschoond.

6. **Eindrapport** (in chat output):
   - Starttijd, eindtijd, duur
   - Aantal issues opgepakt / afgerond / geblokkeerd
   - Lijst PR-links
   - Eventuele blokkers met reden

## Belangrijke regels

- **Nooit force-push naar main**.
- **Nooit `--no-verify`** — pre-commit hooks moeten slagen.
- **Eén issue tegelijk** — geen parallelle branches.
- **Bij twijfel over scope**: comment op het issue en zet terug op `Todo`. Niet zelf scope-creepen.
- **Acceptance criteria niet gehaald** = niet completed markeren, ook al lijkt het "bijna goed".
- **Geen secrets committen**: scan diff op `.env`, API keys, certs voor je pusht.
