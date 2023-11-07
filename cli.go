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

func PromptUserNewAccount(AccountEntries map[int]*Account) {
	var name string
	var typeAccount AssetType
	var balance decimal.Decimal
	var ID int
	for {
		fmt.Printf("New Account - Enter the name of the new account: ")
		name = ScanName()
		fmt.Printf("New Account - Enter the opening balance for this account: ")
		balance = ScanDollars()
		typeAccount = AssetType(PromptUserForNumber([]string{"Asset", "Liability", "Capital", "Drawing", "Revenue", "Expense"}, "Select Type: ") - 1)
		fmt.Printf("New Account - Enter an ID, please make sure it is unique: ")
		fmt.Scan(&ID)
		if appendAccount(AccountEntries, ID, name, balance, typeAccount) == 0 {
			fmt.Printf("    Success. Account has been added.\n")
			break
		} else {
			fmt.Printf("    Sorry something is not right. Maybe try a different ID...\n")
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
		for {
			fmt.Printf("%d - Enter Account ID: ", count)
			fmt.Scan(&id)
			if _, exist := AccountEntries[id]; !exist {
				fmt.Printf("Account does not exist. Retry...\n")
				//return -1
			} else {
				fmt.Printf("Account found. Proceeding...\n")
				break
			}

		}
		fmt.Printf(" %d - Account #%d - Name: %s | Enter Debit/Credit: ", count, id, AccountEntries[id].Name)
		money = ScanDollars()
		if money.IsZero() {
			fmt.Printf("Empty transaction not counted. Done.\n")
			break
		}
		transaction.Modified[id] = money         // Temporary structure for filling up...
		*Journal = append(*Journal, transaction) // Actual recording of transaction.
		sortJournalByDate(*Journal)              // Make sure dates are ascending...
		if PromptUserForNumber([]string{"Another", "Done"}, "Add another?") == 2 {
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
