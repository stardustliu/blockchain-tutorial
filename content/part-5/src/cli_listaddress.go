package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
)

func (cli *CLI) listAddress() {
	wallets, err := NewWallets()
	if err != nil {
		log.Panic(err)
	}

	addresses := wallets.GetAddress()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
