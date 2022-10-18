package main
import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/2468f3c54a7b498284b55d91676c0913")

	if err != nil {
		log.Fatal("Failed to connect to Eth Node", err)
	}

	cntxt := context.Background()

	txn, pending, _ := conn.TransactionByHash(cntxt, common.HexToHash("0x031140df5a9550adfa9be3fd4b71433ccbabbc6c301040ecdd17964faf185555"))

	if pending {
		fmt.Println("Transaction is pending", txn)
	} else {
		fmt.Println("Transaction is not pending", txn)
	}

}