# Golang 1RM

A lightweight Go package that provides functions to estimate a one-rep max (1RM) using multiple formulas and to predict repetitions based on a given 1RM.

---

## Features

- **Multiple 1RM Formulas**: Epley, Brzycki, Lombardi, Mayhew, Wathan
- **Default Formula Selection**: Automatically picks the best formula based on rep range
- **Rep Prediction**: Invert each formula to predict reps at a given weight and known 1RM
- **Compare All**: Functions to compute and compare results from all formulas at once

---

## Installation

```bash
go get github.com/Logan-Dang/golang1rm
```

Then import it in your Go code:

```go
import "github.com/Logan-Dang/golang1rm"
```

## Quickstart

```go
package main

import (
    "fmt"
    "github.com/yourusername/golang1rm"
)

func main() {
    // Example: 100 kg for 5 reps
    epleyMax := golang1rm.Rm1(100, 5, golang1rm.Epley)
    fmt.Printf("Epley 1RM estimate: %.2f kg\n", epleyMax)

    // Use default (auto) formula selection based on reps
    defaultMax := golang1rm.Rm1Default(100, 5)
    fmt.Printf("Default 1RM estimate: %.2f kg\n", defaultMax)

    // Compare all formulas
    allMaxes := golang1rm.Rm1All(100, 5)
    for formula, max := range allMaxes {
        fmt.Printf("%s formula: %.2f kg\n", formula, max)
    }

    // Predict reps for Epley if your 1RM is 130 kg but you only load 100 kg
    epleyReps := golang1rm.RepPredictEpley(130, 100)
    fmt.Printf("Predicted reps at 100 kg if 1RM = 130 kg (Epley): %.1f\n", epleyReps)
}
```

Run it:

```bash
go run main.go
```

## Usage

### 1. Estimating 1rm

- Rm1(weight, reps, formula): Estimate 1RM for a given formula
- Rm1Default(weight, reps): Auto-select formula based on rep count
- Rm1All(weight, reps): Get estimates from all formulas as a map[Rm1Formula]float64

Example:

```go
brzyckiMax := golang1rm.Rm1(85, 3, golang1rm.Brzycki)
fmt.Println("Brzycki 1RM:", brzyckiMax)
```

### 2. Predicting Reps

- RepPredict(rm1, weight, formula): Predict how many reps you could perform at weight, given rm1
- RepPredictDefault(rm1, weight): Uses Epley by default (but you could adapt)
- RepPredictAll(rm1, weight): Get predictions from all formulas

Example:

```go
reps := golang1rm.RepPredict(120, 100, golang1rm.Brzycki)
fmt.Printf("Predicted reps at 100 kg if 1RM = 120 kg (Brzycki): %.1f\n", reps)
```
