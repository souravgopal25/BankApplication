package db

import (
	"context"
	"database/sql"
	"github.com/souravgopal25/BankApplication/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomTransfer(t *testing.T, account1, account2 *Account) Transfer {
	args := CreateTransferParams{Amount: util.RandomMoney(),
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, args.Amount, transfer.Amount)
	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1, account2 := Get2RandomAccounts(t)
	CreateRandomTransfer(t, &account1, &account2)
}

func TestDeleteTransfer(t *testing.T) {
	account1, account2 := Get2RandomAccounts(t)
	transfer1 := CreateRandomTransfer(t, &account1, &account2)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	transfer2, err1 := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err1)
	require.EqualError(t, err1, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestGetTransfer(t *testing.T) {
	account1, account2 := Get2RandomAccounts(t)
	transfer1 := CreateRandomTransfer(t, &account1, &account2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
}

func TestListTransferFromTo(t *testing.T) {
	account1, account2 := Get2RandomAccounts(t)
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t, &account1, &account2)
	}
	args := ListTransfersFromToParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}
	transfers, err := testQueries.ListTransfersFromTo(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestUpdateTransfer(t *testing.T) {
	account1, account2 := Get2RandomAccounts(t)
	transfer1 := CreateRandomTransfer(t, &account1, &account2)
	args := UpdateTransferParams{
		FromAccountID: transfer1.ToAccountID,
		ToAccountID:   transfer1.FromAccountID,
		Amount:        transfer1.Amount + 1000,
		ID:            transfer1.ID,
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, args.Amount, transfer2.Amount)
	require.Equal(t, args.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer2.ToAccountID)

}
