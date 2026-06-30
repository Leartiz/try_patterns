# Cyclomatic complexity

McCabe cyclomatic complexity (CC): number of independent paths through a function.

Rough rule in Go (gocyclo):

```
CC = 1 + (if + for + case + && + || + ...)
```

Higher CC -> harder to read and to test all paths.

## Essential vs accidental

| Type | Meaning |
|------|---------|
| essential | problem needs many branches (protocol, rules) |
| accidental | same behavior, but nested ifs and noise |

Same inputs, same outputs - different CC.

Check locally:

```bash
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
gocyclo .
```

Metric alone does not say which type you have - read the code.

## Part 1: HTTP status

| Function | Style | gocyclo CC |
|----------|-------|------------|
| `classifyHTTPEssential` | flat switch | 11 |
| `classifyHTTPAccidental` | duplicated range checks | 13 |

## Part 2: routing handle

| Function | Style | gocyclo CC |
|----------|-------|------------|
| `handleRoutingBefore` | nested if-else | 9 |
| `handleRoutingAfter` | guards, switch, extract | 7 |
| `routeRequest` | extracted request flow | 3 |

Techniques in `routing_after.go` (one-line comments in code):

- guard clause - early return, flatten nesting
- switch - message type instead of if-chain
- extract function - `routeRequest` for one flow
  