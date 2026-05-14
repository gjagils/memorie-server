-- Initial Memorie schema (Linear GJA-61 datamodel).
--
-- Four core entities — Person, Place, Event, Relationship — plus a
-- derived memories_index. The index is rebuildable from the core
-- entities + Immich data, so DROPping and regenerating it is safe.

-- +goose Up
-- +goose StatementBegin
CREATE TABLE persons (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    name            TEXT NOT NULL,
    immich_face_ids TEXT NOT NULL DEFAULT '[]',  -- JSON array of Immich person/face UUIDs
    birthday        TEXT,                         -- ISO 8601 date (YYYY-MM-DD), nullable
    deathday        TEXT,                         -- ISO 8601 date, nullable; drives in-memoriam filter
    notes           TEXT NOT NULL DEFAULT '',
    created_at      TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at      TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_persons_birthday ON persons(birthday) WHERE birthday IS NOT NULL;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE places (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL,
    meaning     TEXT NOT NULL DEFAULT '',         -- free-form, e.g. "ons huis", "oma's tuin"
    latitude    REAL,
    longitude   REAL,
    radius_m    INTEGER NOT NULL DEFAULT 100,     -- match-radius for GPS-tagged photos
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE events (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    name        TEXT NOT NULL,
    kind        TEXT NOT NULL CHECK (kind IN ('trip', 'milestone', 'birthday', 'recurring')),
    start_date  TEXT NOT NULL,                    -- ISO 8601 date
    end_date    TEXT,                             -- ISO 8601 date, nullable for single-day events
    place_id    INTEGER REFERENCES places(id) ON DELETE SET NULL,
    notes       TEXT NOT NULL DEFAULT '',
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    updated_at  TEXT NOT NULL DEFAULT (datetime('now'))
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_events_start_date ON events(start_date);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE relationships (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    person_a_id INTEGER NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    person_b_id INTEGER NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    kind        TEXT NOT NULL CHECK (kind IN ('parent', 'child', 'sibling', 'partner', 'friend', 'other')),
    created_at  TEXT NOT NULL DEFAULT (datetime('now')),
    CHECK (person_a_id <> person_b_id),
    UNIQUE (person_a_id, person_b_id, kind)
);
-- +goose StatementEnd

-- +goose StatementBegin
-- Derived/denormalized memory cards. Rebuildable from the entities above
-- plus Immich data; treat as a cache, not source of truth.
CREATE TABLE memories_index (
    id                INTEGER PRIMARY KEY AUTOINCREMENT,
    kind              TEXT NOT NULL,              -- 'on_this_day', 'event', 'relationship', etc.
    memory_date       TEXT NOT NULL,              -- ISO 8601 date this card surfaces on
    payload           TEXT NOT NULL DEFAULT '{}', -- JSON card-specific data
    source_asset_ids  TEXT NOT NULL DEFAULT '[]', -- JSON array of Immich asset UUIDs
    person_ids        TEXT NOT NULL DEFAULT '[]', -- JSON array of persons.id
    place_id          INTEGER REFERENCES places(id) ON DELETE SET NULL,
    event_id          INTEGER REFERENCES events(id) ON DELETE SET NULL,
    created_at        TEXT NOT NULL DEFAULT (datetime('now')),
    expires_at        TEXT                        -- optional TTL for cache eviction
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE INDEX idx_memories_date_kind ON memories_index(memory_date, kind);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_memories_date_kind;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS memories_index;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS relationships;
-- +goose StatementEnd

-- +goose StatementBegin
DROP INDEX IF EXISTS idx_events_start_date;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS places;
-- +goose StatementEnd

-- +goose StatementBegin
DROP INDEX IF EXISTS idx_persons_birthday;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS persons;
-- +goose StatementEnd
