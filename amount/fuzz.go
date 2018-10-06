package amount

import (
    "bytes"
    "encoding/binary"
    dcrd_util "github.com/decred/dcrd/dcrutil"
)

func Fuzz(input []byte) {
    buf := bytes.NewBuffer(input)
    a1, err1 := binary.ReadVarint(buf)
    a2, err2 := binary.ReadVarint(buf)
    b1, err3 := binary.ReadVarint(buf)
    b2, err4 := binary.ReadVarint(buf)
    if err1 == nil && err2 == nil && err3 == nil && err4 == nil && a2 != 0 && b2 != 0 {
        a := float64(a1) / float64(a2)
        b := float64(b1) / float64(b2)
        aAmnt, err := dcrd_util.NewAmount(a)
        if err == nil {
            aAmnt.ToUnit(dcrd_util.AmountMicroCoin)
            aAmnt.Format(dcrd_util.AmountMicroCoin)
            aAmnt.ToCoin()
            aAmnt.String()
            aAmnt.MulF64(b)
        }
    }
}
