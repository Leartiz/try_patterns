# C++ struct padding / sizeof

Same field types, different layout -> different `sizeof`.

## Idea

The compiler inserts **padding** bytes so `int` / `double` sit on aligned addresses.
Field order matters. Access sections (`public:`) usually do not.

## Structs in this example

| Struct | Meaning |
|--------|---------|
| `LayoutScattered` | fields in messy order |
| `LayoutScatteredPublic` | same + `public:` per field (size should match scattered) |
| `LayoutGrouped` | small fields grouped first |
| `LayoutGroupedPublic` | same + `public:` (size should match grouped) |
| `LayoutGroupedPack1` | grouped + `#pragma pack(1)` (usually smallest) |

## Run

```bash
g++ main.cpp -o main
./main.exe
```

Compare the five numbers. Expect roughly:

```
Scattered == ScatteredPublic
Grouped   == GroupedPublic
GroupedPack1  <=  Grouped  <=  Scattered
```

Exact values depend on ABI / compiler / `std::string` size.
