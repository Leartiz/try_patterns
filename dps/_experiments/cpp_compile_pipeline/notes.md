# C++ compile pipeline

Minimal example: see what g++ does step by step.

Usual `g++ main.cpp -o main` runs all stages at once.
Here you stop after each stage and look at the intermediate file.

## Stages

```
main.cpp  -E->  main.ii   preprocess (expand includes / macros)
main.ii   -S->  main.s    compile to assembly
main.s     ->   main.o    assemble
main.o     ->   main.exe  link
```

```
source (.cpp)
    |  -E  preprocessor
    v
expanded text (.ii)
    |  -S  compiler
    v
assembly (.s)
    |  as  assembler
    v
object (.o)
    |  link
    v
program (.exe)
```

## Commands

```bash
g++ -E main.cpp -o main.ii
g++ -S main.ii -o main.s
as main.s -o main.o
g++ main.o -o main
./main.exe
```

One-shot (same result, no intermediate files kept):

```bash
g++ main.cpp -o main
./main.exe
```

## What each command does

### 1. `g++ -E main.cpp -o main.ii` - preprocessor

- expands `#include` (pulls in headers like `iostream`)
- expands `#define` / `#ifdef`
- does **not** compile yet - still C++ text, just much larger

`main.ii` is huge because of headers.

### 2. `g++ -S main.ii -o main.s` - compile to assembly

- turns C++ into human-readable asm for your CPU
- no more classes / `cout` syntax - registers, calls, labels

`main.s` is what the compiler produces before machine code.

### 3. `as main.s -o main.o` - assembler

- turns asm into machine code (object file)
- `main.o` is binary, but **not** a full program yet (no libraries linked)

### 4. `g++ main.o -o main` - link

- links your `.o` with the standard library / runtime
- produces the executable (`main` / `main.exe`)

### 5. `./main.exe` - run

Just run the finished program.

## Artifacts

`*.ii`, `*.s`, `*.o`, `*.exe` are gitignored - regenerate with the commands above.

Look at `main.ii` to see expanded headers; 
look at `main.s` for asm of `main`.
