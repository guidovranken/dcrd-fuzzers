package txscript

import (
    dcrd_txscript "github.com/decred/dcrd/txscript"
    dcrd_wire "github.com/decred/dcrd/wire"
    dcrd_chainhash "github.com/decred/dcrd/chaincfg/chainhash"
    dcrd_chaincfg "github.com/decred/dcrd/chaincfg"
    dcrd_util "github.com/decred/dcrd/dcrutil"
)

func dcrd_DisasmString(input []byte) {
    dcrd_txscript.DisasmString(input)
}

func dcrd_VmStep(input []byte) {
	tx := &dcrd_wire.MsgTx{
		Version: 1,
		TxIn: []*dcrd_wire.TxIn{
			{
				PreviousOutPoint: dcrd_wire.OutPoint{
					Hash: dcrd_chainhash.Hash([32]byte{
						0xc9, 0x97, 0xa5, 0xe5,
						0x6e, 0x10, 0x41, 0x02,
						0xfa, 0x20, 0x9c, 0x6a,
						0x85, 0x2d, 0xd9, 0x06,
						0x60, 0xa2, 0x0b, 0x2d,
						0x9c, 0x35, 0x24, 0x23,
						0xed, 0xce, 0x25, 0x85,
						0x7f, 0xcd, 0x37, 0x04,
					}),
					Index: 0,
				},
				SignatureScript: nil,
				Sequence:        4294967295,
			},
		},
		TxOut: []*dcrd_wire.TxOut{{
			Value:    1000000000,
			PkScript: nil,
		}},
		LockTime: 0,
	}
	vm, err := dcrd_txscript.NewEngine(input, tx, 0, 0, 0, nil)
    if err != nil {
        return
    }
	for i := 0; i < len(input); i++ {
        done, err := vm.Step()
		if err != nil {
            break
		}
		if done {
            break
		}
    }
}

func dcrd_ExtractPKScriptAddrs(input []byte) {
    dcrd_txscript.ExtractPkScriptAddrs(dcrd_txscript.DefaultScriptVersion, input, &dcrd_chaincfg.MainNetParams)
}

func Fuzz(input []byte) {
    /* dcrd */
    if dcrd_txscript.IsMultisigSigScript(input) {
        dcrd_txscript.MultisigRedeemScriptFromScriptSig(input)
        /* Crashes 31-08-2018 dcrd_txscript.GetMultisigMandN(input) */
    }
    isMultiSig, err := dcrd_txscript.IsMultisigScript(input)
    if err == nil && isMultiSig == true {
        dcrd_txscript.CalcMultiSigStats(input)
    }
    dcrd_txscript.IsMultisigSigScript(input)
    dcrd_txscript.GetScriptClass(dcrd_txscript.DefaultScriptVersion, input)
    dcrd_DisasmString(input)
    dcrd_VmStep(input)
    /* Crashed (30-08-2018), confirmed fixed 05-09-2019 */ dcrd_ExtractPKScriptAddrs(input)
    dcrd_txscript.PushedData(input)
    dcrd_txscript.ExtractPkScriptAltSigType(input)
    dcrd_txscript.GenerateProvablyPruneableOut(input)
    dcrd_txscript.IsStakeOutput(input)
    dcrd_txscript.GetStakeOutSubclass(input)
    dcrd_txscript.ContainsStakeOpCodes(input)
    dcrd_txscript.GetScriptHashFromP2SHScript(input)
    dcrd_txscript.PayToScriptHashScript(input)

    {
        builder := dcrd_txscript.NewScriptBuilder()
        builder.Reset()
        builder.AddOps(input)
        builder.AddData(input)
        builder.Script()
    }

    /* Crashes 23-09-2018 */
    if false {
        var l1 uint16
        var l2 uint16
        if len(input) > 4 {
            l1 = uint16(input[0])
            l1 <<= 8
            l1 += uint16(input[1])

            l2 = uint16(input[2])
            l2 <<= 8
            l2 += uint16(input[3])

            l1 %= 4096
            l2 %= 4096

            curInput := input[4:]
            if l1 >= uint16(len(curInput)) {
                l1 = uint16(len(curInput) - 1)
            }
            script1 := curInput[:l1]
            curInput = curInput[l1:]
            if len(curInput) > 0 {
                if l2 >= uint16(len(curInput)) {
                    l2 = uint16(len(curInput) - 1)
                }
                script2 := curInput[:l2]
                dcrd_txscript.CalcScriptInfo(script1, script2, true)
            }
        }
    }
    {
        pos := 0
        left := len(input)
        var keys []*dcrd_util.AddressSecpPubKey
        for {
            if left < 1 {
                break
            }
            keylen := 0
            if input[pos] & 1 == 0 {
                keylen = 33
            } else {
                keylen = 65
            }
            pos += 1
            left -= 1
            if left < keylen {
                break
            }
            key := input[pos:pos+keylen]
            pos += keylen
            left -= keylen
            apk, err := dcrd_util.NewAddressSecpPubKey(key, &dcrd_chaincfg.MainNetParams)
            if err != nil {
                break
            }
            keys = append(keys, apk)
        }
        script, err := dcrd_txscript.MultiSigScript(keys, len(keys))
        if err == nil {
            dcrd_txscript.GetMultisigMandN(script)
        }
    }
    {
        addr, err := dcrd_util.DecodeAddress(string(input))
        if err == nil {
            dcrd_txscript.PayToAddrScript(addr)
        }
    }
}
