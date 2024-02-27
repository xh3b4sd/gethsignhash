package gethsignhash

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/go-cmp/cmp"
)

func Test_GethSignHash_HashMessage(t *testing.T) {
	testCases := []struct {
		byt [][]byte
		has string
	}{
		// Case 000
		{
			byt: [][]byte{
				IntToBytes(7),
			},
			has: "0xcbe76b843b0b2d4ee61809d1578b36268b773207965820ae20ebdaf978bf3afc",
		},
		// Case 001
		{
			byt: [][]byte{
				AddressToBytes("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
			},
			has: "0x207cbaf2b538db65f2c33c1aaa3259d01c8caf9ed3d05efa3d4bdce10948cd76",
		},
		// Case 002
		{
			byt: [][]byte{
				IntToBytes(7),
				IntToBytes(1),
			},
			has: "0xb92fe793e135cccdde4f009e010816d9470d47838f10c9f378379c7c6564fac1",
		},
		// Case 003
		{
			byt: [][]byte{
				IntToBytes(7),
				IntToBytes(1),
				AddressToBytes("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
			},
			has: "0x0fcf785806e52fcea9fad680f96f13a8f4c5932ea6e248a2a05687a64190534b",
		},
		// Case 004
		{
			byt: [][]byte{
				IntToBytes(7),
				IntToBytes(1),
				AddressToBytes("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
				IntToBytes(700),
			},
			has: "0xbc02743c0db4e578cea5caaf1a2f6705faa8e02ef58d14b4611555c22f6e4678",
		},
		// Case 005
		{
			byt: [][]byte{
				IntToBytes(7),
				StringToBytes(":"),
				IntToBytes(1),
				StringToBytes(":"),
				AddressToBytes("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
				StringToBytes(":"),
				IntToBytes(700),
			},
			has: "0x4d84bf5afe05dff160c24b8f734d46a6260d39c21d4551bfa406baf9419aff53",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			has := HashMessage(tc.byt...).String()

			if has != tc.has {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.has, has))
			}
		})
	}
}

func Test_GethSignHash_SignHash(t *testing.T) {
	testCases := []struct {
		has string
		pri string
		sig string
	}{
		// Case 000
		{
			has: "0x4d84bf5afe05dff160c24b8f734d46a6260d39c21d4551bfa406baf9419aff53",
			pri: "fe5e321af0d82edab25abed54f841e575953465e42ce7a51e2008f8807c5582b",
			sig: "0x063e956331f6b03a51aef0e7ac4065ab9e7339b943535a0a1458953436478d995ebabbcfe29b6f93c02ee9fda5954a6290469b758e0df0a08c1c47a630cd6d391c",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var has []byte
			{
				has = common.HexToHash(tc.has).Bytes()
			}

			var key *ecdsa.PrivateKey
			{
				key, err = crypto.HexToECDSA(tc.pri)
				if err != nil {
					t.Fatal(err)
				}
			}

			var sig []byte
			{
				sig, err = SignHash(has, key)
				if err != nil {
					t.Fatal(err)
				}
			}

			if string(sig) != tc.sig {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.sig, sig))
			}
		})
	}
}
