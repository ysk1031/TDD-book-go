# TDD Book - Go Implementation

A Go implementation of the money example from Kent Beck's "Test-Driven Development: By Example" book. This project demonstrates Test-Driven Development (TDD) principles by building a multi-currency money system through incremental development guided by tests.

## Overview

This project implements a money arithmetic system that can handle multiple currencies with conversion rates. It follows the classic TDD approach where tests are written first, followed by the minimal implementation to make them pass, and then refactoring for better design.

## Features

- **Multi-currency support**: Handle different currencies (USD, CHF, etc.)
- **Currency conversion**: Convert between currencies using exchange rates
- **Money arithmetic**: Add and multiply money values
- **Expression evaluation**: Support for complex expressions involving multiple currencies
- **Test-driven design**: Comprehensive test suite demonstrating TDD principles

## Project Structure

```
.
├── main.go              # Simple main entry point
├── go.mod              # Go module definition
└── money/              # Core money package
    ├── money.go        # Money type and operations
    ├── bank.go         # Bank with currency conversion
    ├── expression.go   # Expression interface
    ├── sum.go          # Sum expression implementation
    ├── pair.go         # Currency pair for exchange rates
    └── money_test.go   # Comprehensive test suite
```

## Getting Started

### Prerequisites

- Go 1.24.2 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/ysk1031/TDD-book-go.git
cd TDD-book-go
```

2. Run the tests:
```bash
go test ./money
```

3. Run the application:
```bash
go run main.go
```

## Usage Examples

```go
package main

import (
    "fmt"
    "github.com/ysk1031/tdd-book/money"
)

func main() {
    // Create money instances
    fiveDollars := money.NewDollar(5)
    tenFrancs := money.NewFranc(10)
    
    // Simple arithmetic
    tenDollars := fiveDollars.Times(2)
    
    // Currency conversion with bank
    bank := money.NewBank()
    bank.AddRate("CHF", "USD", 2) // 2 CHF = 1 USD
    
    // Mixed currency addition
    sum := fiveDollars.Plus(tenFrancs)
    result := bank.Reduce(sum, "USD") // Results in $10
    
    fmt.Printf("Result: %v\n", result)
}
```

## Key Concepts Demonstrated

### Test-Driven Development
- **Red-Green-Refactor cycle**: Each feature starts with a failing test
- **Incremental development**: Small steps with immediate feedback
- **Design emergence**: Architecture emerges from test requirements

### Design Patterns
- **Expression pattern**: Unified interface for money and operations
- **Visitor pattern**: Bank.Reduce() method for expression evaluation
- **Value object**: Immutable money instances

### Object-Oriented Design
- **Polymorphism**: Different expressions implement the same interface
- **Encapsulation**: Internal state protected with well-defined interfaces
- **Single Responsibility**: Each class has a focused purpose

## Running Tests

The project includes comprehensive tests that demonstrate TDD principles:

```bash
# Run all tests
go test ./money

# Run tests with verbose output
go test -v ./money

# Run specific test
go test -run TestMultiplication ./money
```

Test coverage includes:
- Money multiplication and equality
- Currency handling
- Simple and mixed-currency addition
- Expression reduction
- Bank currency conversion
- Complex expression combinations

## Contributing

This project serves as an educational example of TDD principles. When contributing:

1. Write tests first
2. Implement minimal code to pass tests
3. Refactor for better design
4. Maintain comprehensive test coverage

## License

This project is for educational purposes, implementing examples from Kent Beck's TDD book.

## References

- [Test-Driven Development: By Example](https://www.amazon.com/Test-Driven-Development-Kent-Beck/dp/0321146530) by Kent Beck
- Original examples adapted to Go programming language