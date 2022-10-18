package main

import (
	"context"
    "crypto/ecdsa"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main(){

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/14ddb510721d4a118b0e9e79ebc9df98")
    if err != nil {
        log.Fatal(err)
    }
address = ""
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "0", err
	}

	return balance.String(), nil
}