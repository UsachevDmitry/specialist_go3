package db

import (
	"context"
	"fmt"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println("Before:", account1.Balance, account2.Balance)

	amount := int64(10)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	result, err := store.TransferTx(context.Background(), TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: amount,
	})
	
	if err != nil {
		fmt.Printf("Error during TransferTx: %v\n", err)
		t.Fail()
		return
	}

	fmt.Printf("Transfer result: %v\n", result)
	fmt.Printf("After:       %v, %v\n", account1.Balance, account2.Balance)
	// if account1.Balance != account1.Balance-amount {
	// 	fmt.Println("Balance mismatch for account1")
	// 	t.Fail()
	// }
	// if account2.Balance != account2.Balance+amount {
	// 	fmt.Println("Balance mismatch for account2")
	// 	t.Fail()
	// }

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", result)
}