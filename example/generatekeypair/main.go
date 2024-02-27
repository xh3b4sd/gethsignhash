package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var err error
	var ok bool

	var key *ecdsa.PrivateKey
	{
		key, err = crypto.GenerateKey()
		if err != nil {
			panic(err)
		}
	}

	var pub *ecdsa.PublicKey
	{
		pub, ok = key.Public().(*ecdsa.PublicKey)
		if !ok {
			panic("public key must be *ecdsa.PublicKey")
		}
	}

	{
		fmt.Println("private key", "   ", hexutil.Encode(crypto.FromECDSA(key))[2:])    // omit 0x
		fmt.Println("public key", "    ", hexutil.Encode(crypto.FromECDSAPub(pub))[4:]) // omit 0x04 (compression flag)
		fmt.Println("eth address", "   ", crypto.PubkeyToAddress(*pub).Hex())
	}
}
