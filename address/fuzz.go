package address

import (
    dcrd_util "github.com/decred/dcrd/dcrutil"
	dcrd_ec "github.com/decred/dcrd/dcrec"
    dcrd_chaincfg "github.com/decred/dcrd/chaincfg"
)

func Fuzz(input []byte) {
    /* dcrd */
    {
        addr, err := dcrd_util.DecodeAddress(string(input))
        if err == nil {
            addr.String()
            addr.EncodeAddress()
            addr.ScriptAddress()
            addr.DSA(&dcrd_chaincfg.MainNetParams)
            addr.Net()
        }
    }

    dcrd_util.NewAddressPubKey(input, &dcrd_chaincfg.MainNetParams)
    dcrd_util.NewAddressPubKeyHash(input, &dcrd_chaincfg.MainNetParams, dcrd_ec.STEcdsaSecp256k1)
    dcrd_util.NewAddressPubKeyHash(input, &dcrd_chaincfg.MainNetParams, dcrd_ec.STEd25519)
    dcrd_util.NewAddressPubKeyHash(input, &dcrd_chaincfg.MainNetParams, dcrd_ec.STSchnorrSecp256k1)
    dcrd_util.NewAddressScriptHash(input, &dcrd_chaincfg.MainNetParams)
    dcrd_util.NewAddressScriptHashFromHash(input, &dcrd_chaincfg.MainNetParams)
    dcrd_util.NewAddressSecpPubKey(input, &dcrd_chaincfg.MainNetParams)
    dcrd_util.NewAddressEdwardsPubKey(input, &dcrd_chaincfg.MainNetParams)
    dcrd_util.NewAddressSecSchnorrPubKey(input, &dcrd_chaincfg.MainNetParams)
}
