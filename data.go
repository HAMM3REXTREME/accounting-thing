package main // Data

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

type AssetType int

const (
	Asset AssetType = iota
	Liability
	Capital
	Drawing
	Revenue
	Expense
)

type Account struct {
	Name    string
	Type    AssetType
	Balance int
}

type Transaction struct {
	Modified    map[int]int // Holds <Account ID --> Credit(+ve)/Debit(-ve)>
	Description string
	Date        Date
}
