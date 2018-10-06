package secp256k1

import (
    "bytes"
    dcrd_chainec "github.com/decred/dcrd/chaincfg/chainec"
)

func Fuzz(input []byte) {
    {
        pk, err := dcrd_chainec.Secp256k1.ParsePubKey(input)
        if err == nil {
            pk1 := (dcrd_chainec.PublicKey)(pk).SerializeUncompressed()
            pk2 := (dcrd_chainec.PublicKey)(pk).SerializeCompressed()
            if !bytes.Equal(pk1, input) {
                if !bytes.Equal(pk2, input) {
                    panic("Serialization error")
                }
            }
        }
    }
    {
        priv, pub := dcrd_chainec.Secp256k1.PrivKeyFromBytes(input)
		_, err := dcrd_chainec.Secp256k1.ParsePubKey(pub.SerializeUncompressed())
        if err == nil {
            hash := []byte{0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9}
            r, s, err := dcrd_chainec.Secp256k1.Sign(priv, hash)
            if err != nil {
                sig := dcrd_chainec.Secp256k1.NewSignature(r, s)

                if !dcrd_chainec.Secp256k1.Verify(pub, hash, sig.GetR(), sig.GetS()) {
                    panic("dcrd_chainec.Secp256k1.Verify")
                }

                pub.Serialize()
                serializedKey := priv.Serialize()
                if !bytes.Equal(serializedKey, input) {
                    panic("dcrd_chainec.Secp256k1: key not equal")
                }
            }
        }

    }
    /* Crashes 31-08-2018 dcrd_chainec.Secp256k1.ParseDERSignature(input) */
    /* Crashes 31-08-2018 dcrd_chainec.Secp256k1.ParseSignature(input) */
}
