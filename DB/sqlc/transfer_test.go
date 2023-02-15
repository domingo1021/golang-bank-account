package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/domingo1021/golang-bank-account/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer{
	randomAccountIDs := GetRandomAccountIDs(2)
	fromAccountID := randomAccountIDs[0]
	toAccountID := randomAccountIDs[1]
	fmt.Printf("FID: %d, TID: %d\n", fromAccountID, toAccountID)
	args := CreateTransferParams{
		FromAccountID: fromAccountID,
		ToAccountID: toAccountID,
		Amount: util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	return transfer
}

// CreateTransfer
func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

// GetTransfer
func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

// ListTransfers
func TestListTransfers(t *testing.T) {
	randomAccountIDs := GetRandomAccountIDs(2)
	fromAccountID := randomAccountIDs[0]
	toAccountID := randomAccountIDs[1]

	args := ListTransfersParams{
		FromAccountID: fromAccountID,
		ToAccountID: toAccountID,
		Limit: 5,
		Offset: 5,
	}
	tranfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.LessOrEqual(t, len(tranfers), 5)

	for _, transfer := range tranfers {
		require.Equal(t, args.FromAccountID, transfer.FromAccountID)
		require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	}

}