package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/shopspring/decimal"
)

const colorCell = Underline + Black + BgWhite
const colorHeader = Underline + Black + Bold + BgYellow

// Import Export Modules

func accountInfo2StdOut(AccountEntries map[int]*Account, Journal []Transaction, delim string) {
	// Print a header
	headers := [4]string{"ID", "Account Name", "Type", "Balance"}
	paddings := [4]int{5, 40, 20, 15}
	left2right := make([]string, 4)
	for i, header := range headers {
		fmt.Printf("%s %s %s%s", colorHeader, padded(header, paddings[i]), ResetAll, delim)
	}
	fmt.Println()
	for id, account := range AccountEntries {
		// For each row
		left2right[0] = strconv.Itoa(id)
		left2right[1] = account.Name
		left2right[2] = GetAccountTypeName(account.Type)
		left2right[3] = getTotalBalance(id, Journal, 0, len(Journal)).StringFixedBank(2)

		// Print with formatting
		for i, cell := range left2right {
			fmt.Printf("%s %s %s%s", colorCell, padded(cell, paddings[i]), ResetAll, delim)
		}
		fmt.Println()
	}

}

func journal2StdOut(Journal []Transaction, AccountEntries map[int]*Account, delim string) {
	// This prints each transaction in a journal sequentially.
	// Needs an AccountEntries argument to get names of accounts from IDs stored in journal.

	// Print a header
	headers := [6]string{"Date", "Day", "Particulars", "P.R.", "Debit", "Credit"}
	paddings := [6]int{15, 5, 40, 5, 10, 10}
	for i, header := range headers {
		fmt.Printf("%s %s %s%s", colorHeader, padded(header, paddings[i]), ResetAll, delim)
	}
	fmt.Println()

	var lastDate Date // Keep track of last date in between printing each transaction
	for _, transaction := range Journal {
		var matrix [][]string                         // Our temporary buffer to store each transaction
		numVerticals := len(transaction.Modified) + 1 // Determine the number of rows (height) to allocate
		// Initialize the 2D temp slice
		for i := 0; i <= numVerticals; i++ {
			row := make([]string, 6)
			matrix = append(matrix, row)
		}

		// Filling the buffer with data starts here
		var column = 0 // Start first row from this column

		matrix[column][1] = strconv.Itoa(transaction.Date.Day) // Write date
		if lastDate.Month != transaction.Date.Month {
			matrix[column][0] = GetMonthName(transaction.Date.Month) + " " // Only write this month if last month is different
		}
		lastDate.Month = transaction.Date.Month // Update our new 'last' month
		if lastDate.Year != transaction.Date.Year {
			matrix[column][0] += strconv.Itoa(transaction.Date.Year) // Only write this year if last year is different
		}
		lastDate.Year = transaction.Date.Year // Update our new 'last' year

		// Walk through each transaction's modified accounts
		for id, money := range transaction.Modified {
			matrix[column][3] = strconv.Itoa(id) // Write Account IDs to 3rd place in our row
			if money.GreaterThan(decimal.Zero) { // What is debit/credit anyways?
				matrix[column][4] = money.StringFixedBank(2) // +ve should go in 4th place (debit)
				matrix[column][2] = AccountEntries[id].Name  // Find name in AccountEntries map using id.
			} else {
				matrix[column][5] = money.Abs().StringFixedBank(2)   // -ve should go in 5th place (credit)
				matrix[column][2] = "    " + AccountEntries[id].Name // Credit entries have an indent
			}
			column = column + 1 // Next column for next modified account and its associated info.

		}
		matrix[numVerticals-1][2] = Italic + transaction.Description + NoItalic // Add transaction description to the bottom, after each modified account

		// Print the buffer in this format
		for c := 0; c <= numVerticals; c++ {
			for left2right := 0; left2right < 6; left2right++ {
				fmt.Printf("%s %s %s%s", colorCell, padded(matrix[c][left2right], paddings[left2right]), ResetAll, delim)
			}
			fmt.Printf("\n")
		}

	}

}

func journal2csv(Journal []Transaction, AccountEntries map[int]*Account, filePath string) error {
	// Needs an AccountEntries argument to get names of accounts from IDs stored in journal.
	var delim string = ", "
	// Print a header
	headers := [6]string{"Date", "Day", "Particulars", "P.R.", "Debit", "Credit"}

	// Open the file for writing. Create the file if it doesn't exist, and truncate it if it does.
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close() // Make sure to close the file when done.

	// Create a buffered writer for efficient writing
	writer := bufio.NewWriter(file)

	for _, header := range headers {
		_, writeErr := writer.WriteString(header + delim)
		if writeErr != nil {
			fmt.Println("Error writing to file:", writeErr)
			return writeErr
		}
	}
	_, writeErr := writer.WriteString("\n")
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return writeErr
	}

	var lastDate Date // Keep track of last date in between printing each transaction
	for _, transaction := range Journal {
		var matrix [][]string                         // Our temporary buffer to store each transaction
		numVerticals := len(transaction.Modified) + 1 // Determine the number of rows (height) to allocate
		// Initialize the 2D temp slice
		for i := 0; i <= numVerticals; i++ {
			row := make([]string, 6)
			matrix = append(matrix, row)
		}

		// Filling the buffer with data starts here
		var column = 0 // Start first row from this column

		matrix[column][1] = strconv.Itoa(transaction.Date.Day) // Write date
		if lastDate.Month != transaction.Date.Month {
			matrix[column][0] = GetMonthName(transaction.Date.Month) + " " // Only write this month if last month is different
		}
		lastDate.Month = transaction.Date.Month // Update our new 'last' month
		if lastDate.Year != transaction.Date.Year {
			matrix[column][0] += strconv.Itoa(transaction.Date.Year) // Only write this year if last year is different
		}
		lastDate.Year = transaction.Date.Year // Update our new 'last' year

		// Walk through each transaction's modified accounts
		for id, money := range transaction.Modified {
			matrix[column][3] = strconv.Itoa(id) // Write Account IDs to 3rd place in our row
			if money.GreaterThan(decimal.Zero) { // What is debit/credit anyways?
				matrix[column][4] = money.StringFixedBank(2) // +ve should go in 4th place (debit)
				matrix[column][2] = AccountEntries[id].Name  // Find name in AccountEntries map using id.
			} else {
				matrix[column][5] = money.Abs().StringFixedBank(2)   // -ve should go in 5th place (credit)
				matrix[column][2] = "    " + AccountEntries[id].Name // Credit entries have an indent
			}
			column = column + 1 // Next column for next modified account and its associated info.

		}
		matrix[numVerticals-1][2] = transaction.Description // Add transaction description to the bottom, after each modified account

		// Write to file from temp var
		for c := 0; c <= numVerticals; c++ {
			for left2right := 0; left2right < 6; left2right++ {
				_, writeErr := writer.WriteString(matrix[c][left2right] + delim)
				if writeErr != nil {
					fmt.Println("Error writing to file:", writeErr)
					return writeErr
				}

			}
			_, writeErr := writer.WriteString("\n")
			if writeErr != nil {
				fmt.Println("Error writing to file:", writeErr)
				return writeErr
			}
		}

		// Flush the writer to ensure that data is written to the file
		if flushErr := writer.Flush(); flushErr != nil {
			fmt.Println("Error flushing writer:", flushErr)
			return flushErr
		}

	}
	return nil
}
