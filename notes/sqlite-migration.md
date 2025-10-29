# Upgrading SQLBoiler for SQLite Support

This repository still contains sqlboiler v2.7.x–generated models that target MySQL. Moving to SQLite with the current major (v4) of sqlboiler requires regenerating every model file, because v4 introduces a new module path (`github.com/aarondl/sqlboiler/v4`) and API changes (context-aware query methods, different helper packages, new null types, etc.). The steps below outline the migration process.

## 1. Install the New Tooling

```bash
go install github.com/aarondl/sqlboiler/v4@v4.19.5
go install github.com/aarondl/sqlboiler/v4/drivers/sqlboiler-sqlite3@v4.19.5
```

The sqlite driver uses the pure-Go `modernc.org/sqlite` backend, so no CGO toolchain is required.

## 2. Prepare a SQLite Schema

SQLBoiler introspects a live database, so create a SQLite database that mirrors the legacy MySQL schema. You can:

1. Export the table definitions from the production MySQL instance (`mysqldump --no-data`), edit for SQLite compatibility (PRIMARY KEY definitions inline, `AUTO_INCREMENT` → `AUTOINCREMENT`, remove backticks, etc.), and apply with `sqlite3 blogwind.db < schema.sql`.
2. Alternatively, craft a minimal schema that includes only the tables your Go code touches (`bloggers`, `articles`, `comments`, etc.) if you plan to trim unused features.

Keep this `.db` file (or a reproducible `.sql`) under `notes/` so future regenerations stay repeatable.

## 3. Create an SQLite sqlboiler Config

Add `sqlboiler.sqlite.toml` alongside the existing MySQL config:

```toml
output      = "models"
pkgname     = "models"
wipe        = true
add-global-variants = true

[sqlite3]
  dbname = "path/to/blogwind.db"
  fkmode = "auto"
```

The `wipe` flag lets sqlboiler overwrite the current MySQL models, so commit or back them up before regenerating.

## 4. Regenerate Models

```bash
sqlboiler sqlite3 --no-tests
```

Expect all generated files to switch imports to `github.com/aarondl/sqlboiler/v4/...` and `github.com/aarondl/null/v8`. Update hand-written code (e.g., `tpl/view_models.go`, `main.go`) to use the new context-based APIs—`qm.Where` still exists, but execution helpers now expect a `context.Context` and an explicit `boil.Executor`.

## 5. Adjust Runtime Wiring

- Switch the runtime driver to `modernc.org/sqlite` (e.g., import `_ "modernc.org/sqlite"` and open connections via `sql.Open("sqlite", "file:blogwind.db?_pragma=busy_timeout=5000")`).
- Revisit all raw SQL (`queries.Raw`, etc.) because SQLite uses positional parameters `?` instead of MySQL-style `?` with type-specific coercions.
- Run `go test ./...` to catch any behavioural drift.

Because these steps touch every model file, tackle them on a clean branch and keep a MySQL-capable version tagged for historical reference.
