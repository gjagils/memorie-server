# Memorie — Personas

Vier persona-profielen die geraadpleegd worden bij elke design- en architectuurkeuze in Memorie. Doel: scherp houden voor wie we bouwen en welke keuzes verschillende gebruikersgroepen winnen of verliezen.

## Hoe te gebruiken bij design keuzes

Bij elke design- of architectuurkeuze (scenario's, UI-flows, install-paden, auth-modellen, AI-gebruik, pricing, packaging, etc.):

1. **Loop alle vier de personas langs.**
2. Beschrijf per persona: **wint** dit hen, **schrikt** het ze af, of is het **neutraal**?
3. Maak de **trade-off expliciet**: welke persona's verliezen we met deze keuze, welke winnen we?
4. Documenteer de uiteindelijke keuze + welke persona-pijn we bewust accepteren.

Geen keuze maken = stilzwijgend kiezen voor de gemakkelijkste persona. Dat is bijna altijd de verkeerde.

## Snelle referentie

| Persona | Heeft Immich/NAS? | Drijfveer | Pijngrens |
|---|---|---|---|
| **Marc** | Ja, volledig ingericht | Ideologisch — weg van big tech | Niet 100% air-gapped |
| **Sanne** | Ja, via Portainer-UI | Gemak + Apple-magie nadoen | >1u install |
| **Tom** | Misschien | Probleem fixen, snel resultaat | >10 min geen output |
| **Lisa** | Nee | Volle telefoon + geen abo's | >1 weekend totale setup |

---

## Persona 1: Marc — De bewuste zelfhoster

**Profiel:** 42, software engineer of IT'er. Heeft Synology + Proxmox draaien, leest Hacker News, gebruikt Bitwarden/Vaultwarden, Nextcloud, gaf Google Photos jaren geleden op. Is principieel: data is zijn data, US-jurisdictie is een rode vlag.

**Setup:** Immich draait al, met externe backup naar een tweede NAS. Tailscale, geen Cloudflare voor private services. Heeft een Docker-stack van 15+ services en weet wat een reverse proxy is.

**Wat triggert hem:** "Eindelijk een memories-laag die níet bij Apple/Google ligt en mijn foto's niet kopieert." Hij wil dat Memorie zo dicht mogelijk bij zijn bestaande Immich blijft.

**Waar haakt hij af:**
- Als data Memorie verlaat richting een externe AI-API (ook Anthropic/OpenAI) zonder expliciete keuze.
- Als Memorie een "tweede kopie" van zijn foto-archief maakt op disk.
- Als er telemetry/analytics in zit.
- Als de install meer is dan `docker compose up`.

**Pijngrens:** "Als dit niet 100% air-gapped kan draaien, gebruik ik het niet."

---

## Persona 2: Sanne — De onbewuste zelfhoster

**Profiel:** 38, marketing-achtergrond, technisch handig maar geen developer. Heeft een Synology gekocht omdat een vriend zei "doe dat", draait Plex, Immich, en Home Assistant. Tegelijk: iCloud Photos vol, Gmail, ChatGPT-betaalaccount, Spotify Family. Privacy is "wel belangrijk hoor" maar gemak wint meestal.

**Setup:** Immich naast iCloud Photos (dubbele backup-gevoel). Foto's komen automatisch via Immich-app erop. Ze kijkt nooit in containers, deployt via Portainer-UI.

**Wat triggert haar:** "Iets leuks doen met al die foto's op mijn Synology — verhalen, automatische albums voor de kinderen, terugkijken." Ze wil Apple Memories-achtige magie maar dan op haar eigen ding.

**Waar haakt ze af:**
- Als ze ook nog een aparte database moet aanmaken/migraties moet draaien.
- Als ze Immich opnieuw moet configureren of foto's moet "importeren".
- Als de UI lelijk of traag is — Apple heeft haar verwend.
- Als familie het niet kan installeren met één link.

**Pijngrens:** "Als ik er een uur aan kwijt ben en het werkt nog niet, doe ik het niet."

---

## Persona 3: Tom — De fixer

**Profiel:** 29, vond Memorie via een Reddit-post of YouTube-video. Geen NAS-fanaat, draait wat in Docker op een mini-PC of een gehuurde VPS. Pragmatisch: "lost dit mijn probleem op? Ja → installeren. Nee → volgende."

**Setup:** Heeft mogelijk **geen Immich**. Misschien een map met foto's, misschien Photoprism, misschien gewoon iCloud waar hij vanaf wil. Komt binnen met een specifiek probleem ("ik wil verhalen rond mijn vakantiefoto's") en zoekt het kortste pad ernaartoe.

**Wat triggert hem:** Eén screenshot of demo-video die zegt: "dit werkt op mijn foto's." Hij leest de README, niet de docs.

**Waar haakt hij af:**
- Als er een harde Immich-dependency is en hij die nog moet opzetten ("nog een ding installeren").
- Als de eerste 5 minuten geen "wow" oplevert.
- Als er meer dan 2 dependencies/containers zijn.

**Pijngrens:** "Als ik na 10 minuten geen output zie, ben ik weg."

---

## Persona 4: Lisa — De volle-telefoon googler

**Profiel:** 31, jonge ouder, twee kinderen onder de 5. Marketeer of HR-rol. Niet technisch maar handig genoeg om een NAS uit de doos te halen als ze moet. iPhone 256GB die alsnog vol zit met video's van de kinderen. iCloud 50GB al lang vol; krijgt al maanden de "iCloud-opslag bijna vol"-meldingen. Wil principieel niet die €2,99/mnd voor altijd betalen.

**Setup vandaag:** Geen NAS, geen Immich, geen Docker-kennis. Misschien iCloud op gratis-tier, foto's via WhatsApp doorgestuurd als noodbackup, vage gedachte "ik moet ooit eens een externe schijf kopen."

**Wat triggert haar:** Googelt "iPhone foto's eigen server" of "iCloud alternatief goedkoop". Komt uit op een blog of Reddit-thread met "draai je eigen Photos-vervanger op een Synology". Memorie is in haar gedachte een **bijvangst** — ze wil van die meldingen af, en dat het ook nog "leuk iets doet met foto's van de kinderen" is een dikke plus.

**Waar haakt ze af:**
- Als ze eerst een NAS moet kopen + Immich apart moet installeren + Memorie als 3e laag erbovenop. Drie hordes voor "ik wil van die meldingen af."
- Als de Memorie-app niet ook de offload-flow regelt (foto's vanaf iPhone → eigen server).
- Als de totale eerste-keer-cost (NAS + tijd) > 5 jaar iCloud-abonnement.

**Pijngrens:** "Als ik er meer dan één weekend in moet steken voordat het werkt, koop ik gewoon iCloud 200GB."

---

## Spanningsvelden tussen de personas

**As 1: heeft al een Immich/NAS-stack?**
- **Ja** (Marc, Sanne) → integratie met Immich is een **feature**, "geen tweede kopie" is verkoopargument.
- **Nee** (Tom, Lisa) → een Immich-dependency is een **drempel**, eventueel een dealbreaker.

**As 2: drijfveer**
- **Ideologisch** (Marc) → wint bij elke "geen big tech / geen externe API"-keuze, verliest bij elk gemak-shortcut dat data uitlekt.
- **Pragmatisch + esthetisch** (Sanne, Lisa) → wint bij "het werkt out-of-the-box" en "het ziet er goed uit", verliest bij elke extra config-stap.
- **Pragmatisch + functioneel** (Tom) → wint bij "snel resultaat", verliest bij elke wachttijd of dependency.

**Kernvraag die deze 4 stellen aan elke design-keuze:**
> Bouwen we een laag bovenop bestaand Immich (Marc/Sanne winnen, Tom/Lisa verliezen)?
> Of een bundle die Immich-installatie meeneemt (Tom/Lisa winnen, Marc vindt het rommelig)?
> Of twee instap-paden die later samenkomen (meer werk, breder bereik)?

Maak die keuze nooit impliciet.
