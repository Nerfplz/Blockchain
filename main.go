package main

import (
	"blockchain/block"
	"blockchain/data"
	"fmt"
)

func main() {
	bc := block.NewBlockchain(data.String("Initial BLOCK"))
	bc.AddBlock(data.String("Some random String Data"))
	bc.AddBlock(data.String(""))
	fmt.Println(bc)
}
