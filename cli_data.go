package main

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
