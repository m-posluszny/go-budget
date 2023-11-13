package accounts

type Account struct {
	Uid       string
	Name      string
	Balance   float64
	Offbudget bool
}

type AccountForm struct {
	Name      string  `form:"name" binding:"required" url:"name"`
	Offbudget bool    `form:"offbudget" binding:"required" url:"offbudget"`
	Initial   float64 `form:"initial" binding:"required" url:"initial"`
}

func GetAccount(accUid string) (Account, error) {
	return Account{Uid: accUid, Name: "Afdasfasd", Balance: 30000, Offbudget: false}, nil
}

func GetAccounts(uid string) ([]Account, []Account, error) {
	var accs []Account
	accs = append(accs, Account{"UID", "A", 100, false})
	accs = append(accs, Account{"UID", "Afdasfasd", 30000, false})
	accs = append(accs, Account{"UID", "Afasvxzc", 100000, false})
	accs = append(accs, Account{"UID", "Afdas", 100, false})
	accs = append(accs, Account{"UID", "fdasfasdA", 100, true})
	accs = append(accs, Account{"UID", "A", 100, true})
	accs = append(accs, Account{"UID", "A", 100, true})
	return accs, accs, nil
}
