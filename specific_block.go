package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "encoding/json"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Transaction data structure
    type Transaction struct {
        Hash     string `json:"hash"`
        Value    string `json:"value"`
        Gas      uint64 `json:"gas"`
        GasPrice uint64 `json:"gas_price"`
        Nonce    uint64 `json:"nonce"`
        To       string `json:"to"`
        Pending  bool   `json:"pending"`
    }
    type NFTActionStack struct {
		BlockNumber int64 `json:"block_number"`
        Timestamp uint64 `json:"timestamp"`
		Difficulty uint64 `json:"difficulty"`
		HeaderNumber string `json:"header_number"`
		Hash string `json:"hash"`
		TransactionsCount int `json:"transaction_count"`
		Transactions []Transaction `json:"transactions"`
	}
	
	// type Connection struct {
	// 	NFTActionStacks []NFTActionStack
	// }

    client, err := ethclient.Dial("https://mainnet.infura.io/v3/2468f3c54a7b498284b55d91676c0913")
    if err != nil {
        log.Fatal(err)
    }

    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Println(header.Number.String()) // 5671744

    blockNumber := big.NewInt(5671744)
    block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }

    // fmt.Println(block.Number().Uint64())     // 5671744 
    // // fmt.Println(block.Time().Uint64())       // 1527211625
    // fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
    // fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
    // fmt.Println(len(block.Transactions()))   // 144

    count, err := client.TransactionCount(context.Background(), block.Hash())
    if err != nil {
        log.Fatal(err)
    }

    // fmt.Println(count) // 144

    if count > 1 {
        // var nftActions []NFTActionStack
        nftActions := &NFTActionStack{
            BlockNumber: block.Number().Int64(),
            Timestamp: block.Time(),
            Difficulty: block.Difficulty().Uint64(),
            HeaderNumber: header.Number.String(),
            Hash: block.Hash().Hex(),
            TransactionsCount: len(block.Transactions()),
            Transactions: []Transaction{},
        }

        for _, tx := range block.Transactions() {
            nftActions.Transactions = append(nftActions.Transactions, Transaction{
                Hash:     tx.Hash().String(),
                Value:    tx.Value().String(),
                Gas:      tx.Gas(),
                GasPrice: tx.GasPrice().Uint64(),
                Nonce:    tx.Nonce(),
                To:       tx.To().String(),
            })
        }
        responseJson, _ := json.Marshal(nftActions)
        fmt.Println(string(responseJson))
    }else{
        fmt.Println("There is no block")
    }
		
}