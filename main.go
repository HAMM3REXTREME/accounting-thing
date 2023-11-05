package main

import (
	"fmt"
)

func main() {
	// Main material to be manipulated
	accountEntries := make(map[int]*Account)
	var Journal []Transaction

	for {
		var userPrompt = PromptUserForNumber([]string{"Add Account", "Add Transaction", "Edit Account", "Edit Transaction"}, "What would you like to do?")
		if userPrompt == 1 {
			PromptUserNewAccount(accountEntries)
		} else if userPrompt == 2 {
			PromptUserNewTransaction(accountEntries, &Journal)
		}

		debugPrintAccounts(accountEntries)

		debugPrintJournal(Journal)
		applyTransaction2Account(Journal, accountEntries)
		fmt.Printf("Applied this transaction.\n")

		debugPrintAccounts(accountEntries)
		fmt.Printf("%s", Journal)

		journal2StdOut(Journal, accountEntries, "|")

		if err := journal2csv(Journal, accountEntries, "transactions.csv"); err != nil {
			fmt.Printf("Error writing CSV: %v\n", err)
		} else {
			fmt.Println("CSV file 'transactions.csv' has been created.")
		}

	}
}

// User Input for Dummy account (as many as wanted)
/* 	for {
	//break
	var inputName string
	fmt.Print("Enter the account name (or 'quit' to exit): ")
	fmt.Scan(&inputName)

	if inputName == "quit" {
		break
	}

	fmt.Print("Enter the opening account balance: ")
	var accountBalance int
	accountBalance = ScanDollars()
	//fmt.Scan(&accountBalance)

	appendAccount(accountEntries, newID, inputName, accountBalance, 1)

	newID++
} */
//appendAccount(accountEntries, 101, "Dank", 32849)
//appendAccount(accountEntries, 102, "Meme", 3453)

// Access and print account information
//debugPrintAccounts(accountEntries)

// Create a test transaction with name, description and a map of (account number --> debit/credit value)
//testTransaction := Transaction{
//	Modified:    make(map[int]int), // Input by user
//	Description: "Test Transaction",
//	Date: Date{
//		Year:  2022,
//		Month: Oct,
//		Day:   1,
//	},
//}
//for id, account := range accountEntries {
//	var debit int
//fmt.Printf("%s Enter the monetary value for Account ID %d: ", account, id)
//	fmt.Printf("Enter a transaction value for account #%d with name \"%s\", type %d and balance of $%s : ", id, account.Name, account.Type, cent2Dollar(account.Balance))
//	debit = ScanDollars()
//	testTransaction.Modified[id] = debit
//}

//Journal = append(Journal, testTransaction) // Add this test transaction to our journal
//Journal = append(Journal, testTransaction) // Add this test transaction to our journal
