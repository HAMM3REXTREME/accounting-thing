package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/shopspring/decimal"
)

func ScanName() string {
	// Scans a line. Will count spaces.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("captured:", line)
	return line
}

func PromptUserForNumber(options []string, header string) int {
	// Takes a list of string options to ask along with a header prompt, then returns the number selected.
	if len(options) == 0 {
		return -1
	}
	fmt.Printf("%s\n", header)
	//fmt.Printf("Select an option by typing a number...\n")
	for number, text := range options {
		fmt.Printf("(\033[1m%d\033[m) - %s \n", number+1, text)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a number from the above options: ")
		scanner.Scan()
		input := scanner.Text()
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(options) {
			fmt.Println("Invalid choice. Please enter a valid number.")
			continue
		}
		return choice
	}
}

func ScanDollars() decimal.Decimal {
	// Returns a decimal.Decimal amount. Use for money values.
	var input string
	//fmt.Print("Enter a dollar value: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return decimal.Zero
	}

	dollar, err := decimal.NewFromString(input)
	return dollar

}

func PromptDateInput() Date {
	// Asks user for Date, and returns a Date.
	var userDate Date
	fmt.Printf("Date - Enter Year: ")
	fmt.Scan(&userDate.Year)
	userDate.Month = MonthInt(PromptUserForNumber([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}, "Date - Enter Month: "))
	fmt.Printf("Date - Enter Day: ")
	fmt.Scan(&userDate.Day)
	return userDate
}

func PromptAccountEdit(AccountEntries map[int]*Account) {
	var name string // Name
	//var contraName string       // Name(s) for contra account
	var typeAccount AccountType // Account type
	//var contraID int            // Contra account id(s)
	var ID int // Main Account IDs
	//var contraAccIDs []int      // List of contra account ids.
	//var contraNames []string    // List of names for contra acounts.

	for {
		// Ask for a valid ID
		fmt.Printf("Edit Account - Enter the ID of the account you want to edit: ")
		fmt.Scan(&ID)
		if _, exist := AccountEntries[ID]; !exist {
			fmt.Printf("Account #%d does not exist. please try again...\n", ID)
			//return -1
		} else {
			fmt.Printf("Account #%d found. Proceeding...\n", ID)
			break
		}
	}

	// After getting a valid ID, ask what to do with it
	userOption := PromptUserForNumber([]string{"Change Name", "Change Type", "Add Contra Account"}, "Add a contra account?")
	if userOption == 1 {
		fmt.Printf("Edit Account - Enter a new name for this account")
		name = ScanName()
		AccountEntries[ID] = &Account{Name: name}
		return
	} else if userOption == 2 {
		typeAccount = AccountType(PromptUserForNumber([]string{"Asset", "Liability", "Capital", "Drawing", "Revenue", "Expense"}, "Select Type: ") - 1)
		AccountEntries[ID] = &Account{Type: typeAccount}
		return
	}

	/* 		for { // Keep asking for contra account IDs and Names, appending them to some lists until user asks to stop.
		if PromptUserForNumber([]string{"Yes", "Done"}, "Add a contra account?") == 2 {
			fmt.Printf("Moving on...\n")
			break
		}
		fmt.Printf("New Contra Account - Enter the name of the contra account: ")
		contraName = ScanName()
		fmt.Printf("New Contra Account - Enter an ID, please make sure it is unique: ")
		fmt.Scan(&contraID)
		contraAccIDs = append(contraAccIDs, contraID)
		contraNames = append(contraNames, contraName)
	} */

}

func PromptUserNewAccount(AccountEntries map[int]*Account, Journal *[]Transaction) {
	// This function will modify the arg variables to add an account (and opening entry) based on input.
	// Keep asking until a valid outcome
	for {
		var name string             // Name
		var contraName string       // Name(s) for contra account
		var typeAccount AccountType // Account type
		var balance decimal.Decimal // Account Opening amount
		var contraID int            // Contra account id(s)
		var ID int                  // Main Account IDs
		var contraAccIDs []int      // List of contra account ids.
		var contraNames []string    // List of names for contra acounts.
		transaction := Transaction{ // Temporary transaction structure for filling up
			Modified:    make(map[int]decimal.Decimal), // Only this account --> opening amount
			Description: "Opening amount for this account",
			Date: Date{
				Year:  0000,
				Month: Jan,
				Day:   0,
			},
		}

		fmt.Printf("New Account - Enter the name of the new account: ")
		name = ScanName()
		fmt.Printf("New Account - Enter the opening balance for this account or enter 0: ")
		balance = ScanDollars()
		typeAccount = AccountType(PromptUserForNumber([]string{"Asset", "Liability", "Capital", "Drawing", "Revenue", "Expense"}, "Select Type: ") - 1)
		fmt.Printf("New Account - Enter an ID, please make sure it is unique: ")
		fmt.Scan(&ID)

		for { // Keep asking for contra account IDs and Names, appending them to some lists until user asks to stop.
			if PromptUserForNumber([]string{"Done", "Yes"}, "Add a contra account?") == 1 {
				fmt.Printf("Continuing...\n")
				break
			}
			fmt.Printf("New Contra Account - Enter the name of the contra account: ")
			contraName = ScanName()
			fmt.Printf("New Contra Account - Enter an ID, please make sure it is unique: ")
			fmt.Scan(&contraID)
			contraAccIDs = append(contraAccIDs, contraID)
			contraNames = append(contraNames, contraName)
		}

		if appendAccount(AccountEntries, ID, name, contraAccIDs, typeAccount, contraNames) == 0 { // Try to append an account
			if !balance.IsZero() { // Only add an opening statement to our Journal if it has a non-zero opening amount.
				transaction.Date = PromptDateInput()
				transaction.Modified[ID] = balance
				*Journal = append(*Journal, transaction) // Add our temp opening transaction to our actual Ledger.
				sortJournalByDate(*Journal)
				fmt.Printf("    Successfully added opening amount to the journal...\n")
			}
			fmt.Printf("    Success. Account has been created.\n")
			break
		} else {
			fmt.Printf("    Sorry, Account cannot be added. Maybe try a different ID (main or contra)...\n")
		}

	}

}

func PromptUserEditTransaction(AccountEntries map[int]*Account, Journal []Transaction) int {
	// This function modifies its args to add a new transaction.
	// Journal is modified, AccountEntries is just used to check if an acount exist

	var transactionNum int
	fmt.Printf("Enter transaction number (starts from 0): ")
	fmt.Scan(&transactionNum)
	if len(Journal) < transactionNum {
		fmt.Printf("This transaction does not exist.\n")
		return -1
	}

			// Temporary structure for filling up
			transaction := Transaction{
				Modified:    Journal[transactionNum].Modified, // Selected transaction
				Description: "foo",
				Date: Date{
					Year:  0000,
					Month: Jan,
					Day:   0,
				},
			}
	var userPrompt = PromptUserForNumber([]string{"Edit Description", "Add/Edit Modified Accounts", "Edit Date"}, "What would you like to do with this transaction?")
	if userPrompt == 1 {
		fmt.Printf("Enter a new desscription for transaction %s: \n",transactionNum)
		Journal[transactionNum].Description = ScanName()
		return 1
	} else if userPrompt == 2 {
		var count int = 1
		// Keep asking for account IDs and their values
		for {
			var id int
			var money decimal.Decimal
			// Keep asking for an account until a valid one is given
			for {
				fmt.Printf("%d - Enter Account ID: ", count)
				fmt.Scan(&id)
				if _, exist := AccountEntries[id]; !exist {
					fmt.Printf("Account #%d does not exist. please try again...\n", id)
					//return -1
				} else {
					fmt.Printf("Account #%d found. Proceeding...\n", id)
					break
				}
	
			}
			// Ask for debit/credit entry for this (valid) account id
			fmt.Printf(" %d - Account #%d - Name: %s | Enter Debit/Credit: ", count, id, AccountEntries[id].Name)
			money = ScanDollars()
			if money.IsZero() { // Only count this entry if money is not zero and if it is zero, exit without adding it to our temp transaction.
				Journal[transactionNum].Modified = transaction.Modified
				fmt.Printf("Empty transaction not counted. Done.\n")
				break
			}
			transaction.Modified[id] = money // When we have our account info done (and money is not zero), we add it here for now.
	
			// Actual recording of transaction, after confirming this is all.
			if PromptUserForNumber([]string{"Another", "Done"}, "Add another?") == 2 {
				Journal[transactionNum].Modified = transaction.Modified

				break
			}
			count++ // Increment the counter for number of accounts inputted
		}
		return 2
	} else if userPrompt == 3 {
		Journal[transactionNum].Date = PromptDateInput()
		sortJournalByDate(Journal)              // Make sure dates are ascending...
		return 3
	}
	return 0
}

func PromptUserNewTransaction(AccountEntries map[int]*Account, Journal *[]Transaction) int {
	// This function modifies its args to add a new transaction.
	// Journal is modified, AccountEntries is just used to check if an acount exist
	// Returns the number of accounts modified by the transaction

	// Temporary structure for filling up
	transaction := Transaction{
		Modified:    make(map[int]decimal.Decimal), // Input by user
		Description: "No description provided",
		Date: Date{
			Year:  0000,
			Month: Jan,
			Day:   0,
		},
	}

	// Ask some info about transaction
	fmt.Printf("New Transaction - Enter a description: ")
	transaction.Description = ScanName()
	transaction.Date = PromptDateInput()
	var count int = 1
	// Keep asking for account IDs and their values
	for {
		var id int
		var money decimal.Decimal
		// Keep asking for an account until a valid one is given
		for {
			fmt.Printf("%d - Enter Account ID: ", count)
			fmt.Scan(&id)
			if _, exist := AccountEntries[id]; !exist {
				fmt.Printf("Account #%d does not exist. please try again...\n", id)
				//return -1
			} else {
				fmt.Printf("Account #%d found. Proceeding...\n", id)
				break
			}

		}
		// Ask for debit/credit entry for this (valid) account id
		fmt.Printf(" %d - Account #%d - Name: %s | Enter Debit/Credit: ", count, id, AccountEntries[id].Name)
		money = ScanDollars()
		if money.IsZero() { // Only count this entry if money is not zero and if it is zero, exit without adding it to our temp transaction.
			*Journal = append(*Journal, transaction) // Add our temp transaction to our actual Ledger.
			sortJournalByDate(*Journal)
			fmt.Printf("Empty transaction not counted. Done.\n")
			break
		}
		transaction.Modified[id] = money // When we have our account info done (and money is not zero), we add it here for now.

		// Actual recording of transaction, after confirming this is all.
		if PromptUserForNumber([]string{"Another", "Done"}, "Add another?") == 2 {
			*Journal = append(*Journal, transaction) // Add our temp transaction to our actual Ledger.
			sortJournalByDate(*Journal)              // Make sure dates are ascending...
			break
		}
		count++ // Increment the counter for number of accounts inputted
	}
	return count
}
