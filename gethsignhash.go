package gethsignhash

import (
	"crypto/ecdsa"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xh3b4sd/tracer"
)

func AddressToBytes(a string) []byte {
	return common.HexToAddress(a).Bytes()
}

func GethSignHash(byt ...[]byte) common.Hash {
	return hash(mess(byt...))
}

func IntToBytes(i int) []byte {
	return common.LeftPadBytes(big.NewInt(int64(i)).Bytes(), 32)
}

func NewSignature(has []byte, pri *ecdsa.PrivateKey) ([]byte, error) {
	var err error

	var sig []byte
	{
		sig, err = crypto.Sign(has, pri)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	if len(sig) == 65 && (sig[64] == 0 || sig[64] == 1) {
		sig[64] += 27
	}

	return []byte(hexutil.Encode(sig)), nil
}

func StringToBytes(a string) []byte {
	return []byte(a)
}

func hash(mes []byte) common.Hash {
	return crypto.Keccak256Hash(
		[]byte{0x19},
		[]byte("Ethereum Signed Message:"),
		[]byte{0x0A},
		[]byte(strconv.Itoa(len(mes))),
		mes,
	)
}

func mess(byt ...[]byte) []byte {
	return crypto.Keccak256Hash(byt...).Bytes()
}
