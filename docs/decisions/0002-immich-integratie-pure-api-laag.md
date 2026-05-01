# ADR 0002 — Immich-integratie: pure API-laag met PhotoSource-interface

- **Datum:** 2026-05-01
- **Status:** Accepted

## Context

Memorie is gepositioneerd als laag bovenop Immich. De vraag: hoe strak/los is die koppeling? Drie scenario's overwogen:

1. **Pure API-laag** — Memorie praat als HTTP-client met Immich's REST API. Geen file-toegang. Eigen DB voor Memorie-specifieke data.
2. **Bundle** — Memorie shipt Immich (met Postgres, Redis, ML) mee in de compose-stack. Eén install voor wie nog niets heeft.
3. **Optioneel-Immich** — Memorie werkt op een foto-folder zonder Immich; mét Immich krijgt het rijkere metadata (faces, GPS, quality).

## Beslissing

**Scenario 1: pure API-laag** voor v0–v2, met **PhotoSource-interface** in de Go-code zodat een tweede provider (PhotoPrism, filesystem, etc.) later toegevoegd kan worden zonder dat alle hogere lagen breken.

Concreet:
- Memorie praat met Immich via z'n REST API. Geen volume-mount, geen DB-access.
- Foto-rendering: redirect of proxy naar Immich's image-endpoints (definitieve keuze in implementatiefase).
- Auth-koppeling: Memorie-user ↔ Immich-user (ontwerp in latere ADR).
- Code-architectuur: alle Immich-calls achter een `PhotoSource` Go-interface in `internal/photosource/`. v0 heeft één implementatie (`ImmichPhotoSource`); de interface zelf kost ~0 extra engineering nu maar bewaakt de blast-radius bij toekomstige Immich-API-changes.

## Personas-jury

| | Oordeel | Reden |
|---|---|---|
| Marc | Wint sterk | Schone laag, geen file-access, geen tweede kopie van foto-data |
| Sanne | Wint licht | Werkt naast haar bestaande Immich; Memorie is gewoon "extra Portainer-stack" |
| Tom | Verliest | Immich is harde dependency; "nog een ding installeren" |
| Lisa | Verliest hard | "Eerst Immich, dan Memorie" = 2 hordes, ze haakt af |

## Bewust verlies

- **Tom en Lisa zijn niet de v1-doelgroep.** Memorie's positionering richt zich op mensen die de Apple/Google Photos memories-ervaring waarderen maar hun foto's niet bij big tech willen — dat zijn Marc en Sanne, beiden mensen die al Immich hebben. Tom en Lisa zijn waardevolle latere markten via een toekomstige bundle-installer of standalone-modus, niet nu.
- **Performance-kost:** elke memory-card vereist Immich-API-calls. Mitigatie: caching in Memorie's eigen SQLite (denormalized memory-cards in `memories_index` tabel uit datamodel).
- **Afhankelijkheid van Immich's API-stabiliteit.** Mitigatie: PhotoSource-interface beperkt blast-radius bij breaking changes.

## Afgewezen alternatieven

**Bundle (scenario 2):**
Marc verliest hard — heeft al een gefinetunede Immich met externe backup, gaat geen tweede draaien. Vereist 2 install-paden bouwen ("starter bundle" + "headless naast bestaande Immich"), wat ~2× onderhoud kost en versie-matrix-conflicten oplevert. Conflicteert met "single binary" en "5-min-install" uit ADR-0001 — Immich is een fors multi-container stack. AGPL × Immich's licentie zou ook juridisch nagelopen moeten worden. Mogelijk nuttig als optionele *starter-installer* na v1, niet als architectuur-default nu.

**Optioneel-Immich (scenario 3):**
Vereist mini-Immich-functionaliteit (EXIF, dates, basis-clustering) standalone. Engineering-kost is veel te hoog voor v1 — je bouwt feitelijk een tweede foto-management-laag. Conflicteert met de bevestigde strategische keuze "laag bovenop Immich, niet standalone". Test-matrix verdubbelt (met-Immich-flow vs zonder-Immich-flow). Te overwegen post-v2 als markt erom vraagt; de PhotoSource-interface uit deze ADR houdt die deur open zonder nu iets te bouwen.

## Consequenties

- v0 en v1 hebben Immich als pre-requisite; expliciet noteren in install-docs.
- Code-laag: `internal/photosource/` met `PhotoSource` interface + `ImmichPhotoSource` impl.
- Data-flow: **Immich = source of truth voor foto-metadata**, **Memorie = source of truth voor curatie/relaties/memories-index**.
- README v0: expliciet "vereist Immich"; aparte sectie voor Marc/Sanne over "hoe Memorie naast je bestaande Immich draait".
- Auth-design (Memorie-user ↔ Immich-user): aparte ADR zodra we daar in implementatie tegenaan lopen.
- Caching-strategie voor memory-cards: aparte ADR zodra performance een issue wordt.
