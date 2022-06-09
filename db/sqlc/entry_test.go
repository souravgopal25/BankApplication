package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/souravgopal25/BankApplication/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func CreateRandomEntry(t *testing.T, account *Account) Entry {

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}
func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	CreateRandomEntry(t, &account)
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := CreateRandomEntry(t, &account)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry1, err1 := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err1)
	require.EqualError(t, err1, sql.ErrNoRows.Error())
	require.Empty(t, entry1)
}

func TestGetEntriesOFAccount(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t, &account)
	}
	args := GetEntriesOfAccountParams{
		AccountID: account.ID,
		Offset:    5,
		Limit:     5,
	}
	entries, err := testQueries.GetEntriesOfAccount(context.Background(), args)
	require.NoError(t, err)
	fmt.Println(len(entries))
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := CreateRandomEntry(t, &account)
	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, entry.Amount, entry1.Amount)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.WithinDuration(t, entry.CreatedAt, entry1.CreatedAt, time.Second)
}
func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := CreateRandomEntry(t, &account)
	args := UpdateEntryParams{
		AccountID: entry.AccountID,
		Amount:    entry.Amount + 1000,
		ID:        entry.ID,
	}
	entry1, err := testQueries.UpdateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, args.Amount, entry1.Amount)
	require.Equal(t, args.AccountID, entry1.AccountID)
}
