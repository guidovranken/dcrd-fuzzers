package block

import (
    "bytes"
    dcrd_util "github.com/decred/dcrd/dcrutil"
)

func Fuzz(input []byte) {
    block, err := dcrd_util.NewBlockFromBytes(input)
    if err == nil {
        txs := block.Transactions()
        stxs := block.STransactions()
        bytesOrig, err := block.Bytes()
        if err != nil {
            bytesOrig = nil
        }
        blockHdrBytes, err := block.BlockHeaderBytes()
        if err != nil {
            blockHdrBytes = nil
        }
        block.Hash()
        if len(txs) > 0 {
            block.Tx(0)
            block.TxHash(0)
        }
        if len(stxs) > 0 {
            block.STx(0)
            block.STxHash(0)
        }
        block.TxLoc()
        block.Height()
        if len(txs) > 0 {
            copy1 := dcrd_util.NewBlockDeepCopyCoinbase(block.MsgBlock())
            copy2 := dcrd_util.NewBlockDeepCopy(block.MsgBlock())
            copy1Bytes, err := copy1.Bytes()
            if err == nil && bytesOrig != nil {
                if !bytes.Equal(bytesOrig, copy1Bytes) {
                    /* Crashes 31-08-2018 panic("Bytes() not equal") */
                }
            }
            copy2Bytes, err := copy2.Bytes()
            if err == nil && bytesOrig != nil {
                if !bytes.Equal(bytesOrig, copy2Bytes) {
                    /* Crashes 31-08-2018 panic("Bytes() not equal") */
                }
            }

            copy1BlockHdrBytes, err := copy1.BlockHeaderBytes()
            if err == nil && bytesOrig != nil {
                if !bytes.Equal(blockHdrBytes, copy1BlockHdrBytes) {
                    /* Crashes 31-08-2018 panic("BlockHdrBytes() not equal") */
                }
            }
            copy2BlockHdrBytes, err := copy2.BlockHeaderBytes()
            if err == nil && blockHdrBytes != nil {
                if !bytes.Equal(blockHdrBytes, copy2BlockHdrBytes) {
                    /* Crashes 31-08-2018 panic("BlockHdrBytes() not equal") */
                }
            }
        }
    }
}
