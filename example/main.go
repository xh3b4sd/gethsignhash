package main

import (
	"fmt"

	"github.com/xh3b4sd/gethsignhash"
)

func main() {
	fmt.Println(gethsignhash.GethSignHash(
		gethsignhash.IntToBytes(7),
		gethsignhash.StringToBytes(":"),
		gethsignhash.IntToBytes(1),
		gethsignhash.StringToBytes(":"),
		gethsignhash.AddressToBytes("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
		gethsignhash.StringToBytes(":"),
		gethsignhash.IntToBytes(700),
	).String())
}
