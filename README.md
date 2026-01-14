# Go vs Rust Benchmark

- [Go vs Rust Benchmark](#go-vs-rust-benchmark)
  - [Execution Environment](#execution-environment)
  - [Benchmark Results Summary](#benchmark-results-summary)
  - [Details and Execution Steps](#details-and-execution-steps)
    - [1. CPU (Computation)](#1-cpu-computation)
    - [2. Memory (Allocation)](#2-memory-allocation)
    - [3. Concurrency (Spawning)](#3-concurrency-spawning)

A benchmark comparison of basic performance between Go and Rust.
Tests were conducted across three categories: CPU computation, memory allocation, and concurrency.

## Execution Environment

* **OS**: macOS (darwin)
* **Go Version**: (Based on the execution environment)
* **Rust Version**: (Based on the execution environment)

## Benchmark Results Summary

|    Category     | Language |       Task Details       |    Time / Result     |
| :-------------: | :------: | :----------------------: | :------------------: |
|     **CPU**     |    Go    |   Sum of 100M integers   |      509.29 ms       |
|                 |   Rust   |   Sum of 100M integers   |       72.50 ms       |
|   **Memory**    |    Go    | 10M allocations (struct) | 1.26 s (NumGC: 3168) |
|                 |   Rust   |  10M allocations (Box)   |       10.98 ms       |
| **Concurrency** |    Go    | Spawn 100,000 goroutines |       28.88 ms       |
|                 |   Rust   |   Spawn 10,000 threads   |      963.63 ms       |

> **Note**: In the Concurrency test, Go spawns 100,000 Goroutines. However, since Rust uses OS threads, the count was limited to 10,000 threads to account for resource constraints. Please view this as a demonstration of the architectural differences (Green Threads vs. OS Threads) rather than a direct comparison.

## Details and Execution Steps

### 1. CPU (Computation)

Creates an array of 100 million integers and calculates their sum.

**Go**:

```bash
cd cpu
go run computation.go

```

*Result*: 509.286875ms

**Rust**:

```bash
cd cpu
rustc -C opt-level=3 computation.rs
./computation

```

*Result*: 72.503375ms

### 2. Memory (Allocation)

Performs 10 million object allocations. While Go is impacted by Garbage Collection (GC), Rust deallocates memory immediately via its ownership model.

**Go**:

```bash
cd memory
go run memory.go

```

*Result*: 1.257279s (NumGC = 3168)

**Rust**:

```bash
cd memory
rustc -C opt-level=3 memory.rs
./memory

```

*Result*: 10.976625ms

### 3. Concurrency (Spawning)

Measures the cost of spawning a large number of concurrent tasks.

**Go**:

```bash
cd concurrency
go run concurrency.go

```

*Result*: 28.877ms (100,000 goroutines)

**Rust**:

```bash
cd concurrency
rustc -C opt-level=3 concurrency.rs
./concurrency

```

*Result*: 963.625667ms (10,000 threads)
