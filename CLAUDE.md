# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`github.com/alex-vit/util` — a stdlib-only shared utility library used by other projects in the monorepo via `replace` directives. No external dependencies.

## Build & Test

```bash
go build ./...
go test ./...
go test -run TestName ./ds/   # single test, specifying package
```

## Package Structure

### `package util` (root)

General-purpose helpers:

- **`must.go`** — `Must[V](v, err)` panics on error; `OrExit[T](t, err)` prints and exits.
- **`list_ops.go`** — generic slice operations: `Map`, `Filter`, `FilterInPlace`, `Last`, `LastN`, `Clip`, `RemoveLast`, `RemoveLastN`.
- **`strings.go`** — `Lines` (splits on `\n`), `ContainsIgnoringCase`, `TakeStr`.
- **`aoc.go`** — Advent of Code helpers: `Read` (reads `input.txt`), `Blocks`, `Num`, `Numbers`, `Transposed`, `LCM`, `GCD`, `Abs`. The deprecated `Grid`/`GridStr` functions have been superseded by `ds/grid`.
- **`debug.go`** — `IsDebug` flag (set via `DEBUG` env var), `Debug(s)` conditional print, `SetLogTextOnly`.
- **`files.go`** — `Touch(name)` creates a file if absent.

### `package ds` (`ds/`)

Data structures:

- **`SparseSet[V]`** — int-keyed associative container with O(1) insert/lookup/delete. Uses a sparse index array + dense `Swapback` for cache-friendly iteration via `Entries()`/`Values()`. Keys are non-negative ints (entity IDs). `Get` returns `(V, bool)`.
- **`Swapback[E]`** — unordered slice with O(1) delete-by-index via swap-with-last. Used internally by `SparseSet`. `Delete` returns the index of the element that moved.
- **`Set[V]`** — generic hash set (`map[V]struct{}`). `Add`/`Put` are aliases; both return `bool` indicating if the value was newly added.
- **`DefaultDict[K, V]`** — map that auto-inserts a default value on `Get`. Has `NewDefaultDict` (zero value default) and `NewDefaultDictF` (custom factory).
- **`Queue[V]`** — FIFO queue with `Push`/`Next`/`Iter`. `PQ[V]` is a priority queue using a comparator function (re-sorts on each push — not heap-based).
- **`IntHeap` / `MaxHeap`** — implements `container/heap` interface for min/max int priority queues.
- **`BitSet64`** — uint64-backed bitset with `Set`/`Has`/`Clear`/`Toggle`.
- **`Tup[A, B]`** — generic two-element tuple with destructuring via `.D()`.

### `package grid` (`ds/grid/`)

Grid helper for 2D byte-grid puzzles (primarily AoC). `Grid` stores `[][]byte` with bounds-checked accessors. Direction is encoded as int 0–3 (up/right/down/left). Key methods: `Parse`, `Size`, `Clone`, `Walk`, `Find`, `Is`, `Place`, `Direction` (iterator), `Neighbors`, `Coord`. Panics on non-ASCII input or ragged rows.

## Design Notes

- `SparseSet` uses `Swapback` internally; when `Swapback.Delete` moves the last element, `SparseSet` updates the moved element's sparse index to keep the invariant consistent.
- `Queue.Add` is deprecated in favor of `Queue.Push` — both exist for backward compatibility.
- `Grid.Isb` is deprecated in favor of `Grid.Is` (which accepts a string of acceptable bytes).
- When changing APIs here, update all dependent projects (`game`, `swed`, etc.) in the same session.
