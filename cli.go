package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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

func ScanDollars() int {
	var input string
	//fmt.Print("Enter a dollar value: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 0
	}

	// Remove dollar signs and other non-numeric characters
	RealInput := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || r == '.' || r == '-' {
			return r
		}
		return -1
	}, input)

	// Parse the legit input as a float // Floating Point innacuracies can be ignored for now
	//fmt.Printf("THINGY: %s\n", RealInput)
	value, err := strconv.ParseFloat(RealInput, 64)
	if err != nil {
		fmt.Println("Error parsing input:", err)
		return 0
	}

	// Convert the dollar value to cents
	cents := int(value * 100)

	return cents
}

func PromptUserNewAccount(AccountEntries map[int]*Account) {
	var name string
	var typeAccount AssetType
	var balance int
	var ID int
	for {
		fmt.Printf("New Account - Enter the name of the new account: ")
		fmt.Scanln(&name)
		fmt.Printf("New Account - Enter the opening balance for this account: ")
		balance = ScanDollars()
		typeAccount = AssetType(PromptUserForNumber([]string{"Asset", "Liability", "Capital", "Drawing", "Revenue", "Expense"}, "Select Type: "))
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
	transaction := Transaction{
		Modified:    make(map[int]int), // Input by user
		Description: "No description provided",
		Date: Date{
			Year:  0000,
			Month: Jan,
			Day:   0,
		},
	}

	fmt.Printf("New Transaction - Enter a description: ")
	fmt.Scanln(&transaction.Description)
	transaction.Date = PromptDateInput()
	var count int = 0
	for {
		var id int
		var money int
		count++
		fmt.Printf("%d - Enter Account ID: ", count)
		fmt.Scan(&id)
		if _, exist := AccountEntries[id]; !exist {
			fmt.Printf("Account does not exist. Exiting...\n")
			return -1
		}
		fmt.Printf(" %d - Account #%d - Name: %s | Enter Debit/Credit: ", count, id, AccountEntries[id].Name)
		fmt.Scan(&money)
		transaction.Modified[id] = money
		*Journal = append(*Journal, transaction)
		if PromptUserForNumber([]string{"Another", "Done"}, "Add another?") == 2 {
			break
		}
		return count
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
