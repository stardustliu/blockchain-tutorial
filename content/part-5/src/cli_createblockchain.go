package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

func (cli *CLI) createBlockchain(address string) {
	if !ValidateAdress(address) {
		log.Panic("Error: Address is not valid")
	}

	bc := CreateBlockchain(address)
	bc.db.Close()
	fmt.Println("Done")
}
