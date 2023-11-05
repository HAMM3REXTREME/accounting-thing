package main

import (
	"fmt"
)

func main() {
	// Main material to be manipulated
	accountEntries := make(map[int]*Account)
	var Journal []Transaction

	for {
		// This runs continuously
		debugPrintAccounts(accountEntries)
		// Ask user for options
		var userPrompt = PromptUserForNumber([]string{"Add Account", "Add Transaction", "Edit Account", "Edit Transaction"}, "What would you like to do?")
		if userPrompt == 1 {
			PromptUserNewAccount(accountEntries)
		} else if userPrompt == 2 {
			PromptUserNewTransaction(accountEntries, &Journal)
			applyTransaction2Account(Journal, accountEntries)
			fmt.Printf("Applied this transaction.\n")
		}

		debugPrintAccounts(accountEntries)
		debugPrintJournal(Journal)

		fmt.Println("Journal: ")
		journal2StdOut(Journal, accountEntries, "|")
		fmt.Println()
		fmt.Println("Accounts: ")
		accountInfo2StdOut(accountEntries, "|")
		fmt.Println()

		/* 		if err := journal2csv(Journal, accountEntries, "transactions.csv"); err != nil {
		   			fmt.Printf("Error writing CSV: %v\n", err)
		   		} else {
		   			fmt.Println("CSV file 'transactions.csv' has been created.")
		   		} */

	}
}
