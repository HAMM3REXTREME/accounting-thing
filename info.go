package main

import (
	"fmt"
)

func GetAssetTypeName(assetType AssetType) string {
	names := map[AssetType]string{
		Asset:     "Asset",
		Liability: "Liability",
		Capital:   "Capital",
		Drawing:   "Drawing",
		Revenue:   "Revenue",
		Expense:   "Expense",
	}

	return names[assetType]
}

func GetMonthName(month MonthInt) string {
	names := map[MonthInt]string{
		Jan: "January",
		Feb: "February",
		Mar: "March",
		Apr: "April",
		May: "May",
		Jun: "June",
		Jul: "July",
		Aug: "August",
		Sep: "September",
		Oct: "October",
		Nov: "November",
		Dec: "December",
	}

	return names[month]
}

func debugPrintJournal(Journal []Transaction) {
	for i, t := range Journal {
		fmt.Printf("\033[2mDEBUG: Journal - Transaction %d | Description: %s | Accounts Modified:\n", i, t.Date, t.Description)
		for id, debit := range t.Modified {
			fmt.Printf("    Pending: Account ID: %d | Debit/Credit Entry: %s\n", id, debit.StringFixedBank(2))
		}
	}
	fmt.Printf("\033[m")
}

func debugPrintAccounts(AccountEntries map[int]*Account) {
	for id, account := range AccountEntries {
		fmt.Printf("\033[2mDEBUG: accountEntries has an entry: #%d --> Account(Name: %s,Balance: %s, Type: %d)...\033[m\n", id, account.Name, account.Balance.StringFixedBank(2), account.Type)
		//fmt.Printf("accountEntriesMap: Account Name: %s | Account ID: %d | Balance: %d\n", account.Name, id, account.Balance)
	}
}
