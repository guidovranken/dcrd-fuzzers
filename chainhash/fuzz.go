package chainhash

import (
    dcrd_chainhash "github.com/decred/dcrd/chaincfg/chainhash"
)

func Fuzz(input []byte) {
    dcrd_chainhash.NewHashFromStr(string(input))
}

