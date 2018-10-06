package tx

import (
    dcrd_util "github.com/decred/dcrd/dcrutil"
    dcrd_txsort "github.com/decred/dcrd/dcrutil/txsort"
)

func Fuzz(input []byte) {
    tx, err := dcrd_util.NewTxFromBytes(input)
    if err == nil {
        dcrd_util.NewTx(tx.MsgTx())
        dcrd_util.NewTxDeep(tx.MsgTx())
        dcrd_util.NewTxDeepTxIns(tx.MsgTx())
        msgTx := tx.MsgTx()
        dcrd_txsort.Sort(msgTx)
		dcrd_txsort.InPlaceSort(msgTx)
    }
}
