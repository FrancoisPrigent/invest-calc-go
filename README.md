# Investment Growth Calculator

A containerized CLI tool written in Go to calculate monthly compound interest growth with support for deferred tax calculations. This tool simulates real-world banking by rounding to two decimal places and provides a clear breakdown of "Brut" (Gross) vs. "Net" (After-Tax) returns.

## 🚀 Quick Start

You do not need Go installed locally. You only need **Docker** and **Make**.

```bash
# Build and run with default values
make calc

# Run with custom parameters
# Example: $20k principal, 7.5% rate, 24 months, 15% tax, show history
make calc ARGS="-p 20000 -r 7.5 -d 24 -t 15 -i"

```

## 🛠 Command Line Arguments

| Flag | Description | Default |
| --- | --- | --- |
| `-p` | Principal (Initial Investment amount) | 20000.0 |
| `-f` | Number of months between deposits | 1 |
| `-a` | Amount deposited every `-f` months | 0 |
| `-d` | Duration of the investment in months | 12 |
| `-r` | Annual Interest Rate (e.g., 7.5 for 7.5%) | 7.5 |
| `-t` | Tax Rate on total profit (e.g., 15 for 15%) | 0.0 |
| `-i` | Show month-by-month interest history | false |

## 📊 Features

* **Monthly Compounding:** Calculates interest monthly and reinvests it into the principal.
* **Tax Deferred Model:** Applies tax to the total growth at the end of the duration, maximizing compounding potential.
* **Precision:** Uses `math.Round` to ensure results align with real-world currency standards.
* **Dockerized Environment:** Uses a multi-stage `Go 1.25-alpine` build for maximum performance and minimum image size.
* **Clean CLI:** Optimized Makefile with silent BuildKit execution to hide Docker "noise."

## 

## 🏗 Requirements

* **Docker** (Desktop or Engine)
* **Make** (Build automation tool)

## 📁 Project Structure

* `main.go` — The core Go logic and CLI flag definitions.
* `Dockerfile` — Multi-stage build for Go 1.25.
* `Makefile` — Automation for building, running, and cleaning the container.

---
