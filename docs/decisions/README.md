# Architecture Decision Records (ADR)

Map met alle structurele design- en architectuurkeuzes voor Memorie.

## Hoe werkt dit

Elke significante design-keuze wordt vastgelegd als een ADR (genummerd, gedateerd). Code laat zien *wat* er staat; ADRs leggen vast *waarom* — en welke alternatieven zijn afgewezen.

Elk ADR bevat:
- **Status** — `Accepted`, `Superseded by ADR-XXXX`, of `Deprecated`
- **Personas-jury** — per Marc/Sanne/Tom/Lisa wint/verliest/neutraal (zie [`PERSONAS.md`](../../PERSONAS.md))
- **Bewust verlies** — welke trade-off accepteren we expliciet

## Wanneer een ADR schrijven

**Wel:**
- Stack-keuzes (taal, DB, runtime, infrastructure)
- Architectuur-grenzen (welke component praat met welke, hoe, met welk contract)
- Externe afhankelijkheden (welke service, welke API, welke licentie)
- License-, packaging-, of pricing-keuzes
- Boundary-veranderingen (van "shipt mee" naar "standalone", etc.)

**Niet:**
- Implementatie-details
- Dependency-bumps
- Bug-fixes
- Naming-conventions

## Werkwijze tijdens ontwikkeling

Voor je een architectuurkeuze maakt of een bestaande omgooit:

1. **Lees relevante ADRs** (zoek op woorden uit de titel of doorlees de index).
2. **Loop de 4 personas langs** ([`PERSONAS.md`](../../PERSONAS.md)) — per persona wint/verliest/neutraal.
3. **Schrijf de ADR** met persona-jury én bewust verlies.
4. **Bestaande keuze omgooien?** Nieuwe ADR met `Supersedes ADR-XXXX`, en zet de oude op `Superseded by ADR-YYYY`.
5. **Pas dán** implementeren.

## Naamgeving

`NNNN-korte-titel-met-streepjes.md` waarbij `NNNN` oplopend is (0001, 0002, ...). Eenmaal toegekend nooit hergebruiken, ook niet bij superseded.

## Index

| # | Titel | Datum | Status |
|---|---|---|---|
| [0001](0001-stack-go-sqlite-single-binary.md) | Stack: Go + SQLite + single binary | 2026-05-01 | Accepted |
| [0002](0002-immich-integratie-pure-api-laag.md) | Immich-integratie: pure API-laag | 2026-05-01 | Accepted |
