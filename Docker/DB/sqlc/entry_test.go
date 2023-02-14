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
	entry1 := CreateRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	args := ListEntriesParams{
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}