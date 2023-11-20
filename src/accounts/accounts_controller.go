package accounts

import (
	"github.com/m-posluszny/go-ynab/src/db"
)

func GetAccountFromUid(dbx *db.DBRead, accUid string) (Account, error) {
	var acc Account
	err := dbx.Get(&acc,
		`SELECT * FROM credentials WHERE uid=$1;`,
		accUid)
	return acc, err
}

func GetAccountsFromUserUid(dbx *db.DBRead, q AccountsQuery) ([]Account, error) {
	var accs []Account
	err := dbx.Get(&accs,
		`SELECT * FROM accounts WHERE uid=$1; AND offbudget=$2`,
		q.userUid, q.offBudget)
	return accs, err
}

func CreateAccount(dbx *db.DBWrite, newAcc Account) (*Account, error) {
	_, err := dbx.NamedExec(
		`INSERT INTO accounts (uid, uid, password_hash) VALUES (:username, gen_random_uuid(), :password_hash);`,
		newAcc)
	if err != nil {
		return nil, err
	}

	return &newAcc, nil
}

func DeleteAccount(dbx *db.DBWrite, accUid string) error {
	_, err := dbx.Exec(
		`DELETE FROM accounts WHERE uid=$1;`,
		accUid)
	return err
}
