package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries // embedding
	db       *sql.DB
}

// new a *Query for general use
// for transaction, will use db.BeginTx() create new sql.TX to new a *query (like line 32.)
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	//tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
	}
	return tx.Commit()
}

// TransferTxParams contains the input paramters of the transfer transaction.
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    *Transfer `json:"transfer"`
	FromAccount *Account  `json:"from_account"`
	ToAccount   *Account  `json:"to_account"`
	FromEntry   *Entry    `json:"from_entry"`
	ToEntry     *Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// q *Query is put from execTx transaction.
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		var newTransfer Transfer
		var fromEntry, toEntry Entry
		var bigAccount, smallAccount Account

		// create transfer
		newTransfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}
		result.Transfer = &newTransfer

		// create From Account Entry
		fromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry = &fromEntry

		// create From Account Entry
		toEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry = &toEntry

		// manage account balance in ID ADES order.
		fromAccountBigger := arg.FromAccountID > arg.ToAccountID
		bigAccount, err = q.AddAccountBalance(ctx, AddAccountFactory(fromAccountBigger, arg)[0])
		if err != nil {
			return err
		}
		smallAccount, err = q.AddAccountBalance(ctx, AddAccountFactory(fromAccountBigger, arg)[1])
		if err != nil {
			return err
		}
		result.resultAccountMatch(fromAccountBigger, &bigAccount, &smallAccount)

		return nil
	})

	return result, err
}

func AddAccountFactory(fromAccountBigger bool, arg TransferTxParams) (params [2]AddAccountBalanceParams) {
	var smallIDParams, bigIDParams AddAccountBalanceParams

	if fromAccountBigger {
		bigIDParams = AddAccountBalanceParams{-arg.Amount, arg.FromAccountID}
		smallIDParams = AddAccountBalanceParams{arg.Amount, arg.ToAccountID}
	} else {
		bigIDParams = AddAccountBalanceParams{arg.Amount, arg.ToAccountID}
		smallIDParams = AddAccountBalanceParams{-arg.Amount, arg.FromAccountID}
	}
	params = [2]AddAccountBalanceParams{bigIDParams, smallIDParams}

	return params
}

func (result *TransferTxResult) resultAccountMatch(fromAccountBigger bool, a1 *Account, a2 *Account) {
	if fromAccountBigger {
		result.FromAccount = a1
		result.ToAccount = a2
	} else {
		result.ToAccount = a1
		result.FromAccount = a2
	}
}
