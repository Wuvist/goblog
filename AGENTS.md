# Repository Guidelines

## Project Structure & Module Organization
- `main.go` wires Echo routes and MySQL access; the app reads the `connstr` environment variable at startup.
- `models/` holds sqlboiler-generated ORM structs plus integration-heavy tests; regenerate rather than editing by hand.
- `tpl/` contains view helpers and compiled gorazor templates under `tpl/skins`; author templates in `skins/` and regenerate to refresh code.
- `static/css.go` is the esc-generated asset bundle sourced from `oldweb/Template`; `migrate.sql` tracks schema changes, with `notes/how.md` outlining the production environment.
- `oldweb/` keeps the legacy ASP.NET implementation for reference only—avoid mixing new Go code there.

## Build, Test, and Development Commands
- `go run main.go` starts the API locally; export `connstr="user:pass@tcp(localhost:3306)/blogwind"` (or similar) beforehand.
- `go build -o bin/goblog` produces a deployable binary that embeds templates and static assets.
- `go test ./...` exercises generated model tests and needs a MySQL instance matching `sqlboiler.toml`; use a disposable schema to keep fixtures isolated.
- `sqlboiler mysql` and `gorazor skins tpl/skins` regenerate models and templates after database or view changes; run `esc -o static/css.go -pkg static -include="(.+).css" -prefix="oldweb/" oldweb/Template` when CSS assets change.

## Coding Style & Naming Conventions
- Format Go sources with `gofmt` (or `go fmt ./...`) before committing; keep imports ordered by `goimports`.
- Follow Go naming: exported identifiers use PascalCase, private helpers stay lowerCamelCase, and route handlers end with descriptive verbs (`home`, `blogger`).
- Keep generated directories (`models`, `tpl/skins`, `static/css.go`) pristine; introduce handwritten logic in adjacent packages or helper files.

## Testing Guidelines
- Prefer unit tests for new logic alongside the files under test; integration specs can live next to sqlboiler suites but should guard their own fixtures.
- When running `go test ./models/...`, ensure the MySQL schema is seeded with the data expected by the legacy blogwind tables.
- Record any schema-dependent expectations in `migrate.sql` and describe data requirements in the pull request.

## Commit & Pull Request Guidelines
- Emulate the existing history: short, imperative commit summaries (`Add nginx conf rule`, `Recover avatar`).
- Group related code-generation outputs and their sources in a single commit to ease review.
- PRs should explain the intent, reference related issues, note schema or asset regeneration steps, and attach screenshots or curl traces for user-facing changes when possible.
