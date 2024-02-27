# gethsignhash

This package provides `gethsignhash.NewSignature` to generates signed messages
in Golang that can be verified with [OpenZeppelin.ECDSA] in Solidity.

```go
func NewSignature(has []byte, pri *ecdsa.PrivateKey) ([]byte, error)
```

```go
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
```

```
% go run example/main.go
0x4d84bf5afe05dff160c24b8f734d46a6260d39c21d4551bfa406baf9419aff53
```

[OpenZeppelin.ECDSA]: https://docs.openzeppelin.com/contracts/5.x/api/utils#ECDSA
