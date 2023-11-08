package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func initCLI(accountEntries map[int]*Account, Journal []Transaction) {
	for {
		// This runs continuously
		debugPrintAccounts(accountEntries)
		// Ask user for options
		var userPrompt = PromptUserForNumber([]string{"Add Account", "Add Transaction", "Edit Account", "Edit Transaction"}, "What would you like to do?")
		if userPrompt == 1 {
			PromptUserNewAccount(accountEntries, &Journal)
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

func initGUI() {
	// Create a new builder
	builder, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Error creating builder:", err)
	}

	// Load the Glade XML file
	err = builder.AddFromFile("./main_ui.glade")
	if err != nil {
		log.Fatal("Error loading Glade XML:", err)
	}

	// Create the main window
	obj, err := builder.GetObject("main_window")
	if err != nil {
		log.Fatal("Error getting main_window object:", err)
	}
	window, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal("Error casting main_window object to *gtk.Window")
	}
	// Connect the button signal to an event handler
	button, err := builder.GetObject("journal_button")
	if err != nil {
		log.Fatal("Error getting button object:", err)
	}

	button.(*gtk.Button).Connect("clicked", func() {
		// Replace with actual function to handle this thing
		fmt.Println("Button clicked!")
	})

	// Connect signals and handle user interactions here

	// Show the main window
	window.ShowAll()
}

func main() {
	// Main material to be manipulated
	accountEntries := make(map[int]*Account)
	var Journal []Transaction

	initCLI(accountEntries, Journal)

	//gtk.Init(nil)
	//initGUI() // Initialize GUI
	//gtk.Main() // Start the GTK main loop
}
