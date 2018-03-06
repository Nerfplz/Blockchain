package main

import (
	"fmt"

	"github.com/Nerfplz/blockchain/block"
	"github.com/Nerfplz/blockchain/data"
)

func main() {
	bc := block.NewBlockchain(data.String("Initial BLOCK"))
	bc.AddBlock(data.String("Some random String Data"))
	bc.AddBlock(data.String(""))
	fmt.Println(bc)
}
