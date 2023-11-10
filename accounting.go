package main

import (
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
)

func calculateAccountBalance(AccountEntries map[int]*Account, Journal []Transaction, first int, last int) map[int]decimal.Decimal {
	// Create a map to store account balances
	AccountBalances := make(map[int]decimal.Decimal)

	// Initialize account balances to zero
	for id := range AccountEntries {
		AccountBalances[id] = decimal.NewFromInt(0)
	}
	for i := first; i <= last; i++ {
		for key, value := range Journal[i].Modified {
			AccountBalances[key] = AccountBalances[key].Add(value)
		}

	}
	// Now, AccountBalances contains the balances of each account within the specified range
	return AccountBalances
}

func sortJournalByDate(Journal []Transaction) {
	// This func sorts by Date(int, MonthInt, int), can be improved
	// This is used by sort.Slice()
	less := func(i, j int) bool {
		dateI := Journal[i].Date
		dateJ := Journal[j].Date
		if dateI.Year < dateJ.Year {
			return true
		} else if dateI.Year > dateJ.Year {
			return false
		}
		if dateI.Month < dateJ.Month {
			return true
		} else if dateI.Month > dateJ.Month {
			return false
		}
		return dateI.Day < dateJ.Day
	}

	sort.Slice(Journal, less)
	/*
		 	for num, transaction := range Journal {
				fmt.Printf("DEBUG: New chronological Journal: The %d(st/nd/rd/th) transaction was at %d-%02d-%02d\n", num, transaction.Date.Year, transaction.Date.Month, transaction.Date.Day)
			}
	*/
}

func appendAccount(AccountEntries map[int]*Account, id int, name string, contraAcc []int, typeAccount AccountType, cName []string) int {
	// Look at each associated contra accounts requested to be created before doing anything
	for i, contraID := range contraAcc {
		// Return if requested contra account ID(s) exist
		if _, exist := AccountEntries[contraID]; exist {
			fmt.Printf("Contra Account already exists...\n")
			return -1
		} else {
			// Make contra accounts, keeping the mind the type of contra account
			var contraType AccountType
			if typeAccount == 0 || typeAccount == 2 || typeAccount == 4 {
				contraType = 6
			} else {
				contraType = 7
			}
			AccountEntries[contraID] = &Account{Name: cName[i], Type: contraType, ContraAccounts: []int{}}
		}
	}

	// Return if requested account ID already exists. Do this after contra accounts have been made to avoid confusion
	if _, exist := AccountEntries[id]; exist {
		//fmt.Printf("Account does exist. Exiting...\n")
		return 1
	}

	// Now we add the normal account
	newAccount := &Account{Name: name, Type: typeAccount, ContraAccounts: contraAcc}
	AccountEntries[id] = newAccount
	return 0

}
