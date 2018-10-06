package base58

import (
    dcrd_base58 "github.com/decred/base58"
)

func Fuzz(input []byte) {
    dcrd_base58.Decode(string(input))
    dcrd_base58.Encode(input)
}
