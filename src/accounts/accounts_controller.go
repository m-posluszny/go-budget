package accounts

import (
	"errors"
	"log/slog"
	"time"

	"github.com/m-posluszny/go-ynab/src/db"
)

func GetAccountFromUid(dbx db.Queryable, accUid string) (Account, error) {
	var acc Account
	err := dbx.Get(&acc,
		`SELECT * FROM accounts_balance WHERE uid=$1;`,
		accUid)
	return acc, err
}

func GetAccountsFromUserUid(dbx db.Queryable, q AccountsQuery) ([]Account, error) {
	var accs []Account
	err := dbx.Select(&accs,
		`SELECT * FROM accounts_balance WHERE user_uid::text=$1 AND offbudget=$2;`,
		q.userUid, q.offBudget)
	return accs, err
}

func CreateAccount(dbx db.Queryable, newAcc Account) (string, error) {
	newAcc.CreationDate = time.Now()
	rows, err := dbx.NamedQuery(
		`INSERT INTO accounts (uid, user_uid, name, offbudget, creation_date) VALUES (gen_random_uuid(), :user_uid, :name, :offbudget, :creation_date) RETURNING uid;`,
		newAcc)
	var uid string
	if err != nil {
		slog.Error("E", "create_acc", err)
		return "", err
	}
	if rows.Next() {
		rows.Scan(&uid)
	} else {
		return "", errors.New("no rows returned")
	}
	rows.Close()

	newAcc.Uid = uid

	_, err = dbx.NamedExec(`INSERT INTO transactions (uid,  account_uid, date, payee, memo, value) VALUES (gen_random_uuid(),  :uid, :creation_date, 'Initial', '', :balance);`, newAcc)

	if err != nil {
		slog.Error("E", "create_acc", err)
		return "", err
	}

	return uid, RefreshAll(dbx)
}

func DeleteAccount(dbx db.Queryable, accUid string) error {
	_, err := dbx.Exec(
		`DELETE FROM transactions WHERE account_uid=$1;`,
		accUid)
	if err != nil {
		return err
	}
	_, err = dbx.Exec(
		`DELETE FROM accounts WHERE uid=$1;`,
		accUid)
	if err != nil {
		return err
	}
	return RefreshAll(dbx)
}

func RefreshAll(dbx db.Queryable) error {
	refErr := RefreshPayees(dbx)
	AccErr := RefreshAccountsBalance(dbx)
	transErr := RefreshTransactionView(dbx)
	return errors.Join(refErr, AccErr, transErr)
}

func RefreshPayees(dbx db.Queryable) error {
	_, err := dbx.Exec(
		`REFRESH MATERIALIZED VIEW payees;`)
	return err
}

func RefreshAccountsBalance(dbx db.Queryable) error {
	_, err := dbx.Exec(
		`REFRESH MATERIALIZED VIEW accounts_balance;`)
	return err
}

func RefreshTransactionView(dbx db.Queryable) error {
	_, err := dbx.Exec(
		`REFRESH MATERIALIZED VIEW transactions_view;`)
	return err
}
