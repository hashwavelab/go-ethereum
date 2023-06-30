package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("")
	if err != nil {
		panic(err)
	}

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash("0x..."))
	if err != nil {
		panic(err)
	}

	trace, err := client.TraceCallWithLog(context.Background(), ethereum.CallMsg{
		From:  tx.From(),
		To:    tx.To(),
		Gas:   tx.Gas(),
		Value: tx.Value(),
		Data:  tx.Data(),
	}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(trace)
}
