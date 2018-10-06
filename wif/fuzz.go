package wif

import (
    dcrd_util "github.com/decred/dcrd/dcrutil"
)

func Fuzz(input []byte) {
    wif, err := dcrd_util.DecodeWIF(string(input))
    if err == nil {
        wif.String()
        wif.SerializePubKey()
    }
}
