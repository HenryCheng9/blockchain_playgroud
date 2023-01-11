package main

import "my_blockchain/model"

func main() {
	bc := model.NewBlockChain()

	bc.AddBlock(&model.Header{
		Version: model.DefaultVersion,
		

	})
}