package main
import (
    "context"
    "fmt"
    "log"
    "encoding/json"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {

    type NFTActionStack struct {
		Hash string `json:"hash"`
        Type string `json:"type"`
        Value string `json:"value"`
		Gas uint64 `json:"gas"`
		GasPrice uint64 `json:"gas_price"`
		To string `json:"to"`
        Nonce uint64 `json:"nonce"`
		CollectionName string `json:"collection_name"`
		CollectionAddress string `json:"collection_address"`
	}
	
	// type Connection struct {
	// 	NFTActionStacks []NFTActionStack
	// }

    client, err := ethclient.Dial("https://mainnet.infura.io/v3/2468f3c54a7b498284b55d91676c0913")
    if err != nil {
        log.Fatal(err)
    }

    txHash := common.HexToHash("0xee171aa34742fb3dd1fef90ffa9bde2096c4a250380394ba243be53bc9a387e8")
    tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
    if err != nil {
        log.Fatal(err)
    }

    receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
        if err != nil {
            log.Fatal(err)
        }

    if receipt.Status == 1 {
        if isPending == false {
            // var nftActions []NFTActionStack
            nftActions := &NFTActionStack{
                Hash: tx.Hash().Hex(),
                Type: "mint",
                Value: tx.Value().String(),
                Gas: tx.Gas(),
                GasPrice: tx.GasPrice().Uint64(),
                To: tx.To().String(),
                Nonce: tx.Nonce(),
                CollectionName: "Otherdeed for Otherside",
                CollectionAddress: tx.To().Hex(),
            }
            
            responseJson, _ := json.Marshal(nftActions)
            fmt.Println(string(responseJson))
        }else{
            fmt.Println("this tx haven't be packed")
        }
    }else{
        fmt.Println("Fail")
    }
}