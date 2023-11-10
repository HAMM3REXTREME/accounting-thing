package main // Data

import "github.com/shopspring/decimal" // Use fixed decimal instead of floating point

type MonthInt uint

// Text formatting escape codes
const (
	ResetAll       = "\033[m"
	Bold           = "\033[1m"
	Faint          = "\033[2m"
	Italic         = "\033[3m"
	Underline      = "\033[4m"
	SlowBlink      = "\033[5m"
	RapidBlink     = "\033[6m"
	ReverseVideo   = "\033[7m"
	Conceal        = "\033[8m"
	CrossedOut     = "\033[9m"
	NoBold         = "\033[22m"
	NoItalic       = "\033[23m"
	NoUnderline    = "\033[24m"
	NoBlink        = "\033[25m"
	NoReverseVideo = "\033[27m"
	NoConceal      = "\033[28m"
	NoCrossedOut   = "\033[29m"
)

// Text color escape codes
const (
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Reset   = "\033[39m"
)

// Background color escape codes
const (
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
	BgReset   = "\033[49m"
)

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
	Asset AssetType = iota // I don't get accounting
	Liability
	Capital
	Drawing
	Revenue
	Expense
)

type Account struct {
	Name    string
	Type    AssetType
	Balance decimal.Decimal // Fixed point decimal library
}

type Transaction struct {
	Modified    map[int]decimal.Decimal // Holds int Account ID --> Decimal Money(+ve or -ve)
	Description string
	Date        Date
}
