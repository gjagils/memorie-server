# ADR 0004 — Immich auth via shared API key

- **Datum:** 2026-05-02
- **Status:** Accepted

## Context

Volgens ADR-0002 praat Memorie als HTTP-client met Immich's REST API. Daar moet auth bij. Drie modellen overwogen:

- **A. Shared API-key** — Eén Immich API-key per Memorie-installatie, gezet als env var `IMMICH_API_KEY`. Alle Memorie-profielen (papa/mama/dochters) zien dezelfde Immich-data-view (die van de key-owner).
- **B. Per-user API-keys** — Elk Memorie-profiel koppelt z'n eigen Immich-account/API-key. Multi-account-views.
- **C. OAuth / Immich-as-IDP** — Memorie gebruikt Immich's OAuth-flow voor login.

## Beslissing

**Pad A: shared API-key.**

Concreet:
- `IMMICH_API_KEY` als env var op de Memorie-container (al voorzien in `.env.example` en `docker-compose.yml`).
- Memorie's eigen `Person`-profielen (per familielid, uit Linear GJA-61 datamodel) zijn **lokaal in Memorie's DB** en koppelen niet aan Immich-users. Een Memorie-profiel is een persoon-in-de-foto's-en-verhalen, geen Immich-login.
- API-key wordt aangemaakt in Immich UI: Account Settings → API Keys → New API Key (read-scope volstaat voor v1).
- Memorie's PhotoSource-interface (ADR-0002) krijgt `IMMICH_API_KEY` injected bij constructie van `ImmichPhotoSource`. Toekomstige multi-key (pad B) blijft mogelijk door per-call een ander key-context door te geven.

## Personas-jury

| | Oordeel | Reden |
|---|---|---|
| Marc | Wint sterk | Eén key past hoe hij Immich draait (admin-account met alle foto's); geen extra OAuth-config |
| Sanne | Wint | Klikt key-aanmaak in 1 schermpje; geen Immich-login per Memorie-profiel |
| Tom | Wint | Minder configstappen om "iets te zien" |
| Lisa | Wint | Minder dingen om te begrijpen tijdens setup |

## Bewust verlies

- **Geen multi-account-views in v1.** Families waar elke ouder een eigen Immich-login heeft (zeldzaam in self-host-context, maar mogelijk) krijgen één gedeelde view. Acceptabel voor v1: typische family-Immich heeft 1 admin-account met de hele foto-bibliotheek; lid-accounts zijn uitzondering.
- **Geen per-Memorie-profiel-filtering op basis van Immich-permissies.** Filteren op "wat mag deze persoon zien" gebeurt in Memorie's eigen laag (via Person/Relationship-tabellen), niet in Immich. Ook acceptabel — past bij "Memorie = source of truth voor curatie/relaties" uit ADR-0002.

## Afgewezen alternatieven

**B. Per-user API-keys:**
- Vereist UX in Memorie voor "koppel je Immich-account" per profiel. Mogelijk Immich-account aanmaken voor familieleden die er nog geen hebben.
- Vraagt een datamodel-uitbreiding nu (Memorie's `Person` ↔ Immich-user-id).
- Te veel friction voor v1. **Toekomstig pad blijft expliciet open** via PhotoSource-interface (ADR-0002): een latere ADR kan B implementeren door per-call key-context door te geven.

**C. OAuth / Immich-as-IDP:**
- Immich's OAuth-support is op het moment van schrijven beperkt en niet stabiel; integration-risk hoog.
- Meer engineering voor weinig v1-meerwaarde — een gedeelde key dekt de Marc/Sanne-doelgroep.
- Niet permanent dichtgevoerd: PhotoSource-interface kan later een OAuth-implementatie naast `ImmichPhotoSource` krijgen.

## Consequenties

- `ImmichPhotoSource` in `internal/photosource/immich/` (volgende implementatie-stap) leest `IMMICH_API_KEY` uit constructor-config, niet uit globale state.
- Geen DB-tabel voor "Memorie-profiel ↔ Immich-user-mapping" in initial schema (Linear GJA-61).
- README v0 documenteert key-aanmaak-stappen in Immich UI.
- Errors uit Immich (401/403) loggen als "API key invalid/expired" — operator-actie, geen user-impact.
- Toekomst: per-user-keys = aparte ADR; OAuth-flow = aparte ADR; allebei mogelijk zonder de PhotoSource-laag te breken.
