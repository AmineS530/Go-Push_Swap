# 🧠 Go-Push_Swap — Full Blueprint (AUDHD Edition)

> **TL;DR**: Take a list of numbers. Sort them using only two stacks (A and B) and 11 allowed moves. Print the moves. That's it.

---

## ⚡ THE CONCEPT IN ONE PARAGRAPH

You have **Stack A** (starts with your numbers) and **Stack B** (starts empty).  
You can only do **11 moves** to rearrange them.  
Goal: get A fully sorted (smallest on top), B empty.  
The program figures out the **minimum moves** to do that and prints them.

---

## 🗂️ FILE MAP — What Every File Does

```
Go-Push_Swap/
│
├── exec/
│   ├── push-swap.cmd/Main.go   ← Entry point: reads args → solves → prints moves
│   └── checker.cmd/Main.go     ← Entry point: reads args + move list → says OK or KO
│
└── internal/
    ├── stack/
    │   ├── Helpers.go          ← Defines Node and Stack types + constructors
    │   └── Stack.go            ← Push, Pop, Normalize, NewFromSlice
    │
    ├── parser/
    │   └── Parser.go           ← Validates + converts CLI args to []int
    │
    ├── ops/
    │   ├── Pop.go              ← sa, sb, ss  (swap top two)
    │   ├── Push.go             ← pa, pb      (push between stacks)
    │   ├── Rotate.go           ← ra, rb, rr  (top goes to bottom)
    │   └── RevRotate.go        ← rra, rrb, rrr (bottom comes to top)
    │
    └── solver/
        ├── Solver.go           ← Router: picks the right strategy by size
        ├── SmallSort.go        ← Handles n=2, n=3, n=4/5
        ├── solveBig.go         ← Main algorithm for n>5
        ├── Helpers.go          ← IsSorted, cost math, bestCandidate, moveIndexToTop
        └── insertionHelpers.go ← insertBest, rotateToMin
```

---

## 🔩 THE DATA STRUCTURE — Circular Doubly Linked List

> Think of it like a **bracelet** of numbers. Head points to the "top" of the stack.

```
        ┌─────────────────────────────────┐
        ↓                                 │
   [HEAD=3] ⇄ [7] ⇄ [1] ⇄ [9] ─────────┘
        ↑___prev___prev___prev____________┘
```

Every `Node` has:
- `Val`   → the actual number
- `Index` → its **rank** in sorted order (0 = smallest, n-1 = biggest)
- `next`  → pointer forward (down the stack)
- `prev`  → pointer backward (up / toward bottom)

**Why circular?** Rotate ops (`ra`, `rra`) are **O(1)** — just move the `Head` pointer one step. No copying, no shifting.

---

## 🔢 THE INDEX TRICK — Normalize

Before anything else, numbers get **normalized** to their rank:

```
Input:   [42, 7, 100, 3]
Sorted:  [3,  7, 42, 100]
Ranks:   [2,  1,  3,  0]   ← index of each in sorted order
```

**Why?** So the algorithm only compares ranks (0, 1, 2…), not arbitrary integers.  
Makes everything uniform regardless of what numbers the user gives.

---

## 🎮 THE 11 MOVES — Quick Reference

| Move | What it does |
|------|-------------|
| `sa` | Swap top 2 of A |
| `sb` | Swap top 2 of B |
| `ss` | sa + sb at the same time |
| `pa` | Pop top of B → push to A |
| `pb` | Pop top of A → push to B |
| `ra` | Rotate A: top → bottom |
| `rb` | Rotate B: top → bottom |
| `rr` | ra + rb at the same time |
| `rra` | Reverse rotate A: bottom → top |
| `rrb` | Reverse rotate B: bottom → top |
| `rrr` | rra + rrb at the same time |

---

## 🚦 THE SOLVER ROUTER — `Solver.go`

```
Input size?
  ├─ Already sorted?    → return nothing ✅
  ├─ 2 numbers?         → solveTwo
  ├─ 3 numbers?         → solveThree
  ├─ 4 or 5 numbers?   → solveFive
  └─ 6+ numbers?        → solveBig  (the main algorithm)
```

---

## 🔵 SMALL SORTS

### `solveTwo` (2 numbers)
```
[5, 3]  →  if top > bottom: sa  →  [3, 5] ✅
```
One operation max.

### `solveThree` (3 numbers)
Loops until sorted. Each iteration finds the **biggest element** and moves it to where it can't hurt:
```
top > mid AND top > bot  →  ra   (rotate biggest down)
mid > top AND mid > bot  →  rra  (rev-rotate biggest up then down)
else                     →  sa   (swap the top two)
```
At most 2 operations.

### `solveFive` (4–5 numbers)
1. Move rank `0` (smallest) to top of A → `pb` it to B
2. Move rank `1` (second smallest) to top of A → `pb` it to B
3. Now A has 2–3 elements → `solveThree`
4. `pa` twice to bring 0 and 1 back

---

## 🔴 THE BIG ALGORITHM — `solveBig` (6+ numbers)

> This is a **greedy insertion sort** using B as a holding buffer.

### Step 1 — DUMP to B
```
Push everything from A to B until only 3 remain in A.
```
They go in raw, random order.

### Step 2 — SORT THE 3 IN A
```
Use solveThree on the remaining 3.
A is now: [smallest_3, mid_3, largest_3]  (sorted correctly)
```

### Step 3 — INSERT BACK (the clever part)
For **each node in B**, find the cheapest way to slot it back into A:

#### Finding the target slot in A — `findTargetInA`
A is a sorted circular list. We walk it looking for the "gap" where B's node fits:
```
A circular: [2] → [5] → [8] → [2]  (wraps)
Insert 6?   →  gap between 5 and 8  →  slot before 8

Insert 1?   →  wrap-around gap (8 → 2)  →  slot before 2
```

#### Computing cost — `computeCost`
```
costB = rotations needed to bring B node to top of B
costA = rotations needed to bring target slot to top of A

Same direction (both forward OR both backward)?
    → pay max(costB, costA)   [use rr / rrr to do both at once!]

Different directions?
    → pay costB + costA       [must do separately]
```

#### Finding the best candidate — `bestCandidate`
Scan every node in B, compute cost for each, **pick the cheapest**.

#### Executing the insert — `insertBest`
1. Use `rr` or `rrr` for the **simultaneous** rotations (free savings!)
2. Finish remaining B rotations alone
3. Finish remaining A rotations alone
4. `pa` — pop B's node to top of A (it slots right in)

### Step 4 — FINAL ROTATE
After all inserts, A is sorted but minimum might not be on top.
```
Find rank 0 (the minimum).
Rotate forward or backward (whichever is shorter) until it's on top.
```

---

## 🔄 DATA FLOW — End to End

```
CLI args: "3 1 4 1"
    │
    ▼
parser.ParseArgs()
    ├─ split/flatten args
    ├─ parse ints
    └─ check for duplicates → error if any
    │
    ▼
stack.NewFromSlice(nums)
    ├─ Normalize: [3,1,4] → ranks [1,0,2]
    └─ Build circular linked list (reversed insert so head = first arg)
    │
    ▼
solver.Solve(stackA, stackB)
    └─ routes to solveBig / solveThree / etc.
    │
    ▼
[]string of moves: ["pb", "pb", "ra", "pa", "pa"]
    │
    ▼
fmt.Println each move → stdout
```

---

## 🧮 COST MATH — `costToTop`

```go
forward  = pos           // rotations going ra/rb direction
backward = size - pos    // rotations going rra/rrb direction

if forward <= backward → go forward (cheaper or equal)
else                   → go backward
```

---

## ✅ THE CHECKER (`checker.cmd/Main.go`)

A separate binary that **validates** the solver's output:
1. Parses the same numbers
2. Reads moves line-by-line from stdin
3. Applies each move using the same `ops` package
4. At the end: A sorted AND B empty? → `OK`, else → `KO`

Usage:
```bash
./push-swap "3 1 4 2" | ./checker "3 1 4 2"
# → OK
```

---

## 🗺️ DEPENDENCY GRAPH

```
push-swap.cmd ──► parser ──► (stdlib)
      └──────────► solver ──► stack
                     └──────► ops ──► stack

checker.cmd ───► parser
      ├──────────► ops ──► stack
      ├──────────► stack
      └──────────► solver (IsSorted only)
```

---

## 🔑 KEY INSIGHT — Why This Works Efficiently

The algorithm is **greedy**: at every step it picks the globally cheapest insertion.
The `rr`/`rrr` simultaneous rotation trick is what saves the most moves —
when A and B both need rotating in the same direction, you get **two rotations for the price of one**.

This typically achieves ~700 moves for 100 numbers and ~5500 for 500 numbers.
