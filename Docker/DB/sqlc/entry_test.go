package db

import (
	"context"
	"testing"

	"github.com/domingo1021/golang-bank-account/util"
	"github.com/stretchr/testify/require"
)

// create random entry
func CreateRandomEntry(t *testing.T) Entry {
	accountID := GetRandomAccountID()
	randAmount := util.RandomMoney()
	account, err := testQueries.GetAccount(context.Background(), accountID)
	require.NoError(t, err)
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount: randAmount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	return entry
}

//create entry unit test
/*
	args: args.AccountID, args.Amount
*/
func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {

}

func TestListEntries(t *testing.T) {

}