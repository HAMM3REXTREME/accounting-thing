package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/shopspring/decimal"
)

func ScanName() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("captured:", line)
	return line
}

func PromptUserForNumber(options []string, header string) int {
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

func PromptUserNewAccount(AccountEntries map[int]*Account, Journal *[]Transaction) {
	var name string
	var typeAccount AssetType
	var balance decimal.Decimal
	var ID int
	// Temporary structure for filling up
	transaction := Transaction{
		Modified:    make(map[int]decimal.Decimal), // Only this account --> opening amount
		Description: "Opening amount for this account",
		Date: Date{
			Year:  0000,
			Month: Jan,
			Day:   0,
		},
	}
	for {
		fmt.Printf("New Account - Enter the name of the new account: ")
		name = ScanName()
		fmt.Printf("New Account - Enter the opening balance for this account or enter 0: ")
		balance = ScanDollars()
		typeAccount = AssetType(PromptUserForNumber([]string{"Asset", "Liability", "Capital", "Drawing", "Revenue", "Expense"}, "Select Type: ") - 1)
		fmt.Printf("New Account - Enter an ID, please make sure it is unique: ")
		fmt.Scan(&ID)
		if appendAccount(AccountEntries, ID, name, balance, typeAccount) == 0 {
			if !balance.IsZero() { // Only add an opening statement to our Journal if it has a non-zero opening amount.
				transaction.Date = PromptDateInput()
				transaction.Modified[ID] = balance
				*Journal = append(*Journal, transaction) // Add our temp opening transaction to our actual Ledger.
				sortJournalByDate(*Journal)
				fmt.Printf("    Success. Account has been added along with opening amount transaction to journal.\n")
			}
			fmt.Printf("    Success. Account has been added with no opening amount.\n")
			break
		} else {
			fmt.Printf("    Sorry, Account cannot be added. Maybe try a different ID...\n")
		}

	}

}

func PromptUserNewTransaction(AccountEntries map[int]*Account, Journal *[]Transaction) int {
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

	fmt.Printf("New Transaction - Enter a description: ")
	transaction.Description = ScanName()
	transaction.Date = PromptDateInput()
	var count int = 1
	for {
		var id int
		var money decimal.Decimal
		// Keep asking for account IDs and their values
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
		// Ask for debit/credit entry for this account id
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
		count++
	}
	return count
}

func PromptDateInput() Date {
	var userDate Date
	fmt.Printf("Date - Enter Year: ")
	fmt.Scan(&userDate.Year)
	userDate.Month = MonthInt(PromptUserForNumber([]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}, "Date - Enter Month: "))
	fmt.Printf("Date - Enter Day: ")
	fmt.Scan(&userDate.Day)
	return userDate
}
