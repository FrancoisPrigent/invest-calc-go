package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

type Investment struct {
	Principal        float64
	Deposit          Deposit
	AnnualRate       float64
	Duration         int
	TaxRate          float64
}

type Deposit struct {
	Frequence int     // in months, default 1 (monthly)
	Amount    float64
}

// Process handles the logic and populates the result struct.
func (inv Investment) Process() InvestmentResult {
	currentBalance := inv.Principal
	history := make([]float64, inv.Duration)
	var totalGrossGrowth float64

	// Calculate Gross Growth (Compounding).
	for m := 0; m < inv.Duration; m++ {
		// Add monthly deposit at the beginning of the month.
		if (inv.Deposit.Frequence > 0) && ((m%inv.Deposit.Frequence) == 0) {
			currentBalance += inv.Deposit.Amount
		}

		// Calculate monthly interest.
		interest := (currentBalance * inv.AnnualRate / 100) / 12
		interest = math.Round(interest*100) / 100

		// Store interest and update totals.
		history[m] = interest
		totalGrossGrowth += interest
		currentBalance += interest
	}

	// Calculate Tax on the total profit (Deferred Tax).
	taxDue := totalGrossGrowth * (inv.TaxRate / 100)
	taxDue = math.Round(taxDue*100) / 100

	return InvestmentResult{
		Params:         inv,
		GrossGrowth:    totalGrossGrowth,
		GrossFinal:     currentBalance,
		TaxAmount:      taxDue,
		NetGrowth:      totalGrossGrowth - taxDue,
		NetFinal:       currentBalance - taxDue,
		MonthlyHistory: history,
	}
}

type InvestmentResult struct {
	Params         Investment
	GrossGrowth    float64
	GrossFinal     float64
	NetGrowth      float64
	NetFinal       float64
	TaxAmount      float64
	MonthlyHistory []float64
}

// Display handles the formatting of the results.
func (ir InvestmentResult) Display(showHistory bool) {
	fmt.Printf("--- Investment Results ---\n")
	fmt.Printf("Principal:       $%.2f\n", ir.Params.Principal)
	fmt.Printf("Deposit:         $%.2f every %d months\n", ir.Params.Deposit.Amount, ir.Params.Deposit.Frequence)
	fmt.Printf("Annual Rate:     %.2f%%\n", ir.Params.AnnualRate)
	fmt.Printf("Tax Rate:        %.2f%%\n", ir.Params.TaxRate)
	fmt.Printf("Duration:        %d months\n", ir.Params.Duration)
	fmt.Println("--------------------------")

	fmt.Printf("Total investment: $%.2f\n", ir.Params.Principal+float64(ir.Params.Duration/ir.Params.Deposit.Frequence)*ir.Params.Deposit.Amount)

	fmt.Printf("\nBRUT (Gross):\n")
	fmt.Printf("  Total Growth: $%.2f\n", ir.GrossGrowth)
	fmt.Printf("  Final Amount: $%.2f\n", ir.GrossFinal)

	if ir.Params.TaxRate > 0 {
		fmt.Printf("\nNET (After Tax):\n")
		fmt.Printf("  Tax Withheld: $%.2f\n", ir.TaxAmount)
		fmt.Printf("  Total Growth: $%.2f\n", ir.NetGrowth)
		fmt.Printf("  Final Amount: $%.2f\n", ir.NetFinal)
	}

	if showHistory {
		fmt.Printf("\nMonthly Gross Interest History:\n")
		for month, interest := range ir.MonthlyHistory {
			fmt.Printf("Month %2d: $%.2f\n", month+1, interest)
		}
	}
}

// This program calculates the final amount of an investment after a certain number of months
// with a fixed monthly interest rate applied to the cumulative amount.
func main() {
	inv := Investment{}

	flag.Float64Var(&inv.Principal, "p", 20000.0, "Principal (initial investment)")
	flag.IntVar(&inv.Deposit.Frequence, "f", 1, "Number of months between deposits `-a`")
	flag.Float64Var(&inv.Deposit.Amount, "a", 0, "Amount deposited every `-f` months")
	flag.IntVar(&inv.Duration, "d", 12, "Duration in months")
	flag.Float64Var(&inv.AnnualRate, "r", 7.5, "Annual Interest Rate (e.g. 7.5 for 7.5%)")
	flag.Float64Var(&inv.TaxRate, "t", 0, "Tax rate on profit (e.g. 15 for 15%)")
	showHistory := flag.Bool("i", false, "Show interest history")

	flag.Parse()

	// Check for logical errors before running the math.
	if inv.Principal < 0 || inv.Deposit.Frequence < 0 || inv.Deposit.Amount < 0 || inv.AnnualRate < 0 || inv.Duration <= 0 || inv.TaxRate < 0 {
		fmt.Println("Error: Invalid input parameters. Ensure values are positive.")
		os.Exit(1)
	}

	// Calculate and get result struct.
	result := inv.Process()

	// Use the struct's method to print output.
	result.Display(*showHistory)
}
