package main

import (
	"fmt"
	"log"
)

func cliStart(accountEntries map[int]*Account, Journal []Transaction) {
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

func onButtonClicked() {
	log.Println("Button clicked!")
}

func main() {
	// Main material to be manipulated
	accountEntries := make(map[int]*Account)
	var Journal []Transaction
	cliStart(accountEntries, Journal)
	/* gtk.Init(nil)

	builder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Error creating GtkBuilder:", err)
	}

	err = builder.AddFromFile("ui.glade")
	if err != nil {
		log.Fatal("Error loading Glade file:", err)
	}

	win, _ := builder.GetObject("main_window").(*gtk.Window)
	button, _ := builder.GetObject("my_button").(*gtk.Button)

	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Connect the button click event to the onButtonClicked function
	button.Connect("clicked", onButtonClicked)

	win.ShowAll()
	gtk.Main() */
}
