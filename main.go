package main

import (
	"fmt"
	"log"

	"github.com/blck-snwmn/example-go-cov/calculator"
	"github.com/blck-snwmn/example-go-cov/calculator/statistics"
	"github.com/blck-snwmn/example-go-cov/config"
	"github.com/blck-snwmn/example-go-cov/config/env"
	"github.com/blck-snwmn/example-go-cov/fileutil"
	"github.com/blck-snwmn/example-go-cov/fileutil/json"
	"github.com/blck-snwmn/example-go-cov/strutil"
	"github.com/blck-snwmn/example-go-cov/strutil/formatter"
)

func main() {
	// Calculator examples
	sum := calculator.Add(10, 5)
	diff := calculator.Subtract(10, 5)
	product := calculator.Multiply(10, 5)
	quotient, err := calculator.Divide(10, 5)
	if err != nil {
		log.Fatalf("Division error: %v", err)
	}

	// Advanced calculator examples
	pow, err := calculator.Power(2, 3)
	if err != nil {
		log.Fatalf("Power calculation error: %v", err)
	}
	sqrt, err := calculator.Sqrt(16)
	if err != nil {
		log.Fatalf("Square root calculation error: %v", err)
	}

	// Statistics examples
	avg := statistics.CalculateAverage([]float64{1, 2, 3, 4, 5})
	median, err := statistics.CalculateMedian([]float64{1, 2, 3, 4, 5})
	if err != nil {
		log.Fatalf("Median calculation error: %v", err)
	}

	// String utility examples
	upper := strutil.ToUpper("hello")
	lower := strutil.ToLower("WORLD")
	words := strutil.ParseWords("hello world golang example")

	// Formatter examples
	table := formatter.FormatTable([][]string{
		{"Name", "Age"},
		{"Alice", "30"},
		{"Bob", "25"},
	})

	// File utility examples
	content, err := fileutil.ReadFile("example.txt")
	if err != nil {
		// Just log the error, don't exit
		log.Printf("File read error: %v", err)
	}
	err = fileutil.WriteFile("output.txt", "Hello, World!")
	if err != nil {
		log.Printf("File write error: %v", err)
	}

	// JSON examples
	person, err := json.Parse(`{"name":"John","age":30}`)
	if err != nil {
		log.Printf("JSON parse error: %v", err)
	}

	// Config examples
	cfg, err := config.LoadFromFile("config.json")
	if err != nil {
		log.Printf("Config load error: %v", err)
	}
	valid := config.ValidateConfig(cfg)

	// Environment config examples
	envCfg := env.ParseEnvVars("APP_")

	// Display results
	fmt.Printf("Calculator results: sum=%d, diff=%d, product=%d, quotient=%.2f, power=%.2f, sqrt=%.2f\n",
		sum, diff, product, quotient, pow, sqrt)
	fmt.Printf("Statistics results: average=%.2f, median=%.2f\n", avg, median)
	fmt.Printf("String utility results: upper=%s, lower=%s, words=%v\n", upper, lower, words)
	fmt.Printf("Table formatted:\n%s\n", table)
	fmt.Printf("File content: %s\n", content)
	fmt.Printf("JSON parsed: %v\n", person)
	fmt.Printf("Config valid: %v\n", valid)
	fmt.Printf("Environment config: %v\n", envCfg)
}
