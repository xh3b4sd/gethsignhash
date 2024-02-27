package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xh3b4sd/gethsignhash"
)

func main() {
	var err error

	var has []byte
	{
		has = common.HexToHash("0x4d84bf5afe05dff160c24b8f734d46a6260d39c21d4551bfa406baf9419aff53").Bytes()
	}

	var key *ecdsa.PrivateKey
	{
		key, err = crypto.HexToECDSA("fe5e321af0d82edab25abed54f841e575953465e42ce7a51e2008f8807c5582b")
		if err != nil {
			panic(err)
		}
	}

	var sig []byte
	{
		sig, err = gethsignhash.SignHash(has, key)
		if err != nil {
			panic(err)
		}
	}

	{
		fmt.Println(string(sig))
	}
}
