package main

import (
	"fmt"
	"strconv"
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
			fmt.Printf("    Pending: Account ID: %d | Debit/Credit Entry: %d\n", id, debit)
		}
	}
	fmt.Printf("\033[m")
}

func debugPrintAccounts(AccountEntries map[int]*Account) {
	for id, account := range AccountEntries {
		fmt.Printf("\033[2mDEBUG: accountEntries has an entry: #%d --> Account(Name: %s,Balance: %d, Type: %d)...\033[m\n", id, account.Name, account.Balance, account.Type)
		//fmt.Printf("accountEntriesMap: Account Name: %s | Account ID: %d | Balance: %d\n", account.Name, id, account.Balance)
	}
}

func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func cent2Dollar(cents int) string {
	// Convert an int value of cents to a string of dollars (No $$$ sign)
	strCents := strconv.Itoa(cents)
	onlyBucks := string(strCents[:len(strCents)-2]) // 0 to (last-2)
	onlyCents := string(strCents[len(strCents)-2:]) // (last-2) to last
	return onlyBucks + "." + onlyCents              // "999" + "." + "99"
}
