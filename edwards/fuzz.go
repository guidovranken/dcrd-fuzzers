package edwards

import (
    "bytes"
    dcrd_chainec "github.com/decred/dcrd/chaincfg/chainec"
)

func Fuzz(input []byte) {
    dcrd_chainec.Edwards.ParsePubKey(input)
    dcrd_chainec.Edwards.PrivKeyFromBytes(input)
    {
        priv, pub := dcrd_chainec.Edwards.PrivKeyFromScalar(input)
        if priv != nil && pub != nil {
            hash := []byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9}
            r, s, err := dcrd_chainec.Edwards.Sign(priv, hash)
            if err == nil {
                sig := dcrd_chainec.Edwards.NewSignature(r, s)
                if !dcrd_chainec.Edwards.Verify(pub, hash, sig.GetR(), sig.GetS()) {
                    panic("dcrd_chainec.Edwards.Verify")
                }
                pub.Serialize()
                serializedKey := priv.Serialize()
                if !bytes.Equal(serializedKey, input) {
                    panic("dcrd_chainec.Edwards: key not equal")
                }
            }
        }
    }
    dcrd_chainec.Edwards.ParseDERSignature(input)
    dcrd_chainec.Edwards.ParseSignature(input)
}
