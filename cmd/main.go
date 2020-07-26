package main

import (
	"log"
	"os"

	"github.com/oshikawatkm/bitcoin-message/pkg/internal/spv"
	"github.com/oshikawatkm/bitcoin-message/pkg/protocol/network"
)

func main() {

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	c := network.NewClient("10.4.255.146:18333")
	defer c.Conn.Close()
	log.Printf("[+]remote addr: %s", c.Conn.RemoteAddr().String())

	spv := spv.NewSPV(c)
	if err := spv.Handshake(0); err != nil {
		log.Fatal("[-] handshake error; ", err)
	}
	log.Printf("[+] address: %s", spv.Wallet.GetAddress())

	command := os.Args[1]
	switch command {
	case "balance":
		if err := spv.SendFilterLoad(); err != nil {
			log.Fatal("[-] filterload error: ", err)
		}

		if err := spv.SendGetBlocks("000000000000014ad045b835a2f4990a6acedccd95e8e3f42c0fe8caccba05a5"); err != nil {
			log.Fatal("[-] GetBlocks error: ", err)
		}

		if err := spv.MessageHandlerForBalance(); err != nil {
			log.Fatal("[-] main: message handler err: ", err)
		}

		balance := spv.Wallet.GetBanalce()
		log.Printf("[-] Balance: %d", balance)
	case "send":
		log.Printf("[+] send")
		if err := spv.SendGetBlocks("000000000000014ad045b835a2f4990a6acedccd95e8e3f42c0fe8caccba05a5"); err != nil {
			log.Fatal("[-] GetBlocks error: ", err)
		}

		if err := spv.MessagehandlerForBalance(); err != nil {
			log.Fatal("[-] main: message handler err: ", err)
		}
		balance := spv.Wallet.GetBalance()
		log.Printf("[+] Bitcoin: %d", balance)
		transaction := spv.SendTxInv("mgavKSS3hKCAyLKFhy5VHTYu5CMj8AAxQV", 1000)
		if err := spv.MessageHandlerForSend(transaction); err != nil {
			log.Fatal("[+] main: message handler err:", err)
		}
	default:
		log.Printf("[-] no command")
	}

	log.Printf("[*] finish")
}
