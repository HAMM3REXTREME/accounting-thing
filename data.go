package main // Data

import "github.com/shopspring/decimal" // Use fixed decimal instead of floating point

type MonthInt uint

const (
	Jan = 1
	Feb = 2
	Mar = 3
	Apr = 4
	May = 5
	Jun = 6
	Jul = 7
	Aug = 8
	Sep = 9
	Oct = 10
	Nov = 11
	Dec = 12
)

type Date struct {
	Year  int
	Month MonthInt
	Day   int
}

type AccountType int

const (
	Asset AccountType = iota // I don't get accounting
	Liability
	Capital
	Drawing
	Revenue
	Expense
	ContraAsset     // Idk
	ContraLiability // Idk
)

type Account struct {
	Name           string
	Type           AccountType
	ContraAccounts []int
}

type Transaction struct {
	Modified    map[int]decimal.Decimal // Holds int Account ID --> Decimal Money(+ve or -ve)
	Description string
	Date        Date
}
