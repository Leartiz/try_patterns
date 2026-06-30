# Layout

Study repo: theory separate from code per language.
Each example is a self-contained mini-project.

## Basis

```
theory/_basis/<area>/<topic>/notes.md          # theory, language-agnostic
dps_go/_basis/<area>/<topic>/v1/               # canonical code (Go)
dps/_basis/<area>/<topic>/v1/                  # same in C++ (when needed)
dps_py/_basis/<area>/<topic>/v1/               # same in Python (when needed)
```

`<area>` - `data_structures`, `algorithms`, ...

## Design patterns

```
theory/<category>/<pattern>/notes.md           # theory
dps_go/<category>/<pattern>/v1/                # code (Go)
dps/<category>/<pattern>/v1/                   # code (C++)
dps_py/<category>/<pattern>/v1/                # code (Python)
```

`<category>` - `structural`, `behavioral`, `data_source`, `optimization`, ...

## Utility folders (Go)

| Folder | Purpose |
|--------|---------|
| `_basis` | study examples, canonical implementations |
| `_systems` | composed subsystems (auth, http api, migrations, ...) |
| `_experiments` | quick drafts and spikes |

## Go modules

- Each `v1/` has its own `go.mod` - independent mini-project.
- Run: `cd .../v1 && go run .`

## Two repositories

| Repository | Role |
|------------|------|
| `try_patterns` | learn properly: theory + canonical examples |
| `programming_experiments` | sandbox: LeetCode, quick drafts |
