package main

import (
	"sort"

	"github.com/shopspring/decimal"
)

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

func applyTransaction2Account(Journal []Transaction, accountEntries map[int]*Account) {
	for _, t := range Journal {
		for id, debit := range t.Modified {
			updateAccountBalance(accountEntries, id, debit)
			//fmt.Printf("        Applied this transaction.\n")
		}
	}
}

func appendAccount(AccountEntries map[int]*Account, id int, name string, balance decimal.Decimal, typeAccount AssetType) int {
	if _, exist := AccountEntries[id]; exist {
		//fmt.Printf("Account does exist. Exiting...\n")
		return 1
	}
	newAccount := &Account{Name: name, Type: typeAccount, Balance: balance}
	AccountEntries[id] = newAccount
	return 0

}

func updateAccountBalance(AccountEntries map[int]*Account, id int, debit decimal.Decimal) int {
	if _, exist := AccountEntries[id]; !exist {
		//fmt.Printf("Oops: Account does not exist. Exiting...\n")
		return 1
	}
	AccountEntries[id].Balance = decimal.Sum(AccountEntries[id].Balance, debit)
	return 0
}
