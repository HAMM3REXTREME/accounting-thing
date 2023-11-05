package main

func applyTransaction2Account(Journal []Transaction, accountEntries map[int]*Account) {
	for _, t := range Journal {
		for id, debit := range t.Modified {
			updateAccountBalance(accountEntries, id, debit)
			//fmt.Printf("        Applied this transaction.\n")
		}
	}
}

func appendAccount(AccountEntries map[int]*Account, id int, name string, balance int, typeAccount AssetType) int {
	if _, exist := AccountEntries[id]; exist {
		//fmt.Printf("Account does exist. Exiting...\n")
		return 1
	}
	newAccount := &Account{Name: name, Type: typeAccount, Balance: balance}
	AccountEntries[id] = newAccount
	return 0

}

func updateAccountBalance(AccountEntries map[int]*Account, id int, debit int) int {
	if _, exist := AccountEntries[id]; !exist {
		//fmt.Printf("Account does not exist. Exiting...\n")
		return 1
	}
	AccountEntries[id].Balance += debit
	return 0
}
