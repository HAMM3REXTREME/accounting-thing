package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Import Export Modules

func accountInfo2StdOut(AccountEntries map[int]*Account, delim string) {
	// Print a header
	headers := [4]string{"ID", "Account Name", "Type", "Balance"}
	left2right := make([]string, 4)
	for _, header := range headers {
		fmt.Printf("\033[4;1;43m %-20s \033[m%s", header, delim)
	}
	fmt.Println()
	for id, account := range AccountEntries {
		// For each row
		left2right[0] = strconv.Itoa(id)
		left2right[1] = account.Name
		left2right[2] = GetAssetTypeName(account.Type)
		left2right[3] = cent2Dollar(account.Balance)

		// Print with formatting
		for _, cell := range left2right {
			fmt.Printf("\033[4;47m %-20s \033[m%s", cell, delim)
		}
		fmt.Println()
	}

}

func journal2StdOut(Journal []Transaction, AccountEntries map[int]*Account, delim string) {
	// This prints each transaction in a journal sequentially.
	// Needs an AccountEntries argument to get names of accounts from IDs stored in journal.

	// Print a header
	headers := [6]string{"Date", "Day", "Particulars", "P.R.", "Debit", "Credit"}
	for _, header := range headers {
		fmt.Printf("\033[4;1;43m %-20s \033[m%s", header, delim)
	}
	fmt.Println()

	var lastDate Date // Keep track of last date in between printing each transaction
	for _, transaction := range Journal {
		var matrix [][]string                        // Our temporary buffer to store each transaction
		var numVerticals int                         // number of rows (height) to allocate
		numVerticals = len(transaction.Modified) + 1 // Determine the number of rows (height)
		// Initialize the 2D temp slice
		for i := 0; i <= numVerticals; i++ {
			row := make([]string, 6)
			matrix = append(matrix, row)
		}

		// Filling the buffer with data starts here
		var column = 0 // Start first row from this column

		matrix[column][1] = strconv.Itoa(transaction.Date.Day) // Write date
		if lastDate.Month != transaction.Date.Month {
			matrix[column][0] = GetMonthName(transaction.Date.Month) // Only write this month if last month is different
		}
		lastDate.Month = transaction.Date.Month // Update our new 'last' month
		if lastDate.Year != transaction.Date.Year {
			matrix[column][0] += (" " + strconv.Itoa(transaction.Date.Year)) // Only write this year if last year is different
		}
		lastDate.Year = transaction.Date.Year // Update our new 'last' year

		// Walk through each transaction's modified accounts
		for id, money := range transaction.Modified {
			matrix[column][3] = strconv.Itoa(id) // Write Account IDs to 3rd place in our row
			if money > 0 {
				matrix[column][4] = cent2Dollar(money)      // +ve should go in 4th place (debit)
				matrix[column][2] = AccountEntries[id].Name // Find name in AccountEntries map using id.
			} else {
				matrix[column][5] = cent2Dollar(intAbs(money))       // -ve should go in 5th place (credit)
				matrix[column][2] = "    " + AccountEntries[id].Name // Credit entries have an indent
			}
			column = column + 1 // Next column for next modified account and its associated info.

		}
		matrix[numVerticals-1][2] = transaction.Description // Add transaction description to the bottom, after each modified account

		// Print the buffer in this format
		for c := 0; c <= numVerticals; c++ {
			for left2right := 0; left2right < 6; left2right++ {
				fmt.Printf("\033[4;47m %-20s \033[m%s", matrix[c][left2right], delim)
			}
			fmt.Printf("\n")
		}

	}

}

func journal2csv(Journal []Transaction, Accounts map[int]*Account, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// Write the header row
	header := []string{"Date", "Particulars", "P.R.", "Debit", "Credit"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write each transaction to the CSV
	for _, transaction := range Journal {
		for accountID, debit := range transaction.Modified {
			row := []string{fmt.Sprint(transaction.Date.Year), fmt.Sprint(Accounts[accountID].Name), fmt.Sprint(accountID), fmt.Sprint(debit)}
			if err := writer.Write(row); err != nil {
				return err
			}
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
