package main

import "BlockChain/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("Send 1 BTC tp Jacky")
	bc.SendData("Send 1 EOS tp Jacky")
	bc.Print()
}
