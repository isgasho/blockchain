package main

import (
	"encoding/json"
	"fmt"

	"github.com/prologic/blockchain"
)

func main() {
	c := blockchain.NewChain()
	c.Write([]byte("foo"))
	c.Write([]byte("bar"))
	if !c.Verify() {
		panic("blockchain verification failure")
	}

	data, _ := json.Marshal(c)
	fmt.Print(string(data))
}
