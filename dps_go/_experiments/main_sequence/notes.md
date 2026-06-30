# Main sequence (abstractness vs instability)

From package metrics (Martin / architecture books).
Plot **abstractness A** on Y-axis, **instability I** on X-axis.

## Formulas

Abstractness:

```
A = Na / N

Na = number of abstract types (interfaces in Go)
N  = number of all types (interfaces + structs)
```

Instability:

```
I = Ce / (Ca + Ce)

Ca = afferent coupling  (packages that depend on this one)
Ce = efferent coupling  (packages this one depends on)

I = 0  -> stable   (many depend on you, you depend on few)
I = 1  -> unstable (you depend on many, few depend on you)
```

> Чем больше узлов от которых зависит модуль, тем хуже.
> Ce большой: модуль хрупче, сложнее тестировать, сложнее переиспользовать.

Distance from main sequence:

```
D = |A + I - 1|

D = 0  -> on main sequence (balanced)
D = 1  -> farthest from main sequence (bad corner)
```

Main sequence is the line where **A + I = 1**.

Examples:

```
pain corner:        A=0, I=0      ->  A+I=0  ->  D = |0 - 1| = 1  (bad)
balanced point:     A=0.5, I=0.5  ->  A+I=1  ->  D = 0  (good)
useless corner:     A=1, I=1      ->  A+I=2  ->  D = |2 - 1| = 1  (bad)
concrete leaf:      A=0, I=1      ->  A+I=1  ->  D = 0  (good)
```

## Zones

| Zone | A | I | A + I | D | Meaning |
|------|---|---|-------|---|---------|
| pain | low (~0) | low (~0) | ~0 | ~1 | concrete + stable: hard to change |
| uselessness | high (~1) | high (~1) | ~2 | ~1 | abstract + unstable: hard to use |
| main sequence | - | - | ~1 | ~0 | balance abstract vs concrete |

## Packages in this experiment

| Package | Idea |
|---------|------|
| `pain` | one god struct: DB, email, validation inline |
| `useless` | factories and interfaces, awkward to call |
| `balanced` | small interface + one concrete impl |

## Rough types (this repo only)

| Package | interfaces | structs | A (rough) | I (in real project) |
|---------|------------|---------|-----------|---------------------|
| pain | 0 | 1 | 0 | low if many import it |
| useless | 5 | 4 | ~0.56 | depends on coupling |
| balanced | 1 | 2 | ~0.33 | depends on coupling |

This toy repo is too small for real Ca/Ce. 
Use formulas on a real module graph.

To place **balanced** on main sequence, raise I (more dependents) or raise A slightly -
target **A + I near 1**, not **A=0 and I=1** unless you want a concrete leaf.
