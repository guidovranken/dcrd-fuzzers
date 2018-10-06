package wire

import (
    "encoding/binary"
	"bytes"
	"io"

    dcrd_chainhash "github.com/decred/dcrd/chaincfg/chainhash"
    dcrd_wire "github.com/decred/dcrd/wire"
)

// fixedWriter implements the io.Writer interface and intentially allows
// testing of error paths by forcing short writes.
type fixedWriter struct {
	b   []byte
	pos int
}

// Write writes the contents of p to w.  When the contents of p would cause
// the writer to exceed the maximum allowed size of the fixed writer,
// io.ErrShortWrite is returned and the writer is left unchanged.
//
// This satisfies the io.Writer interface.
func (w *fixedWriter) Write(p []byte) (n int, err error) {
	lenp := len(p)
	if w.pos+lenp > cap(w.b) {
		return 0, io.ErrShortWrite
	}
	n = lenp
	w.pos += copy(w.b[w.pos:], p)
	return
}

// Bytes returns the bytes already written to the fixed writer.
func (w *fixedWriter) Bytes() []byte {
	return w.b
}

// newFixedWriter returns a new io.Writer that will error once more bytes than
// the specified max have been written.
func newFixedWriter(max int) io.Writer {
	b := make([]byte, max)
	fw := fixedWriter{b, 0}
	return &fw
}

// fixedReader implements the io.Reader interface and intentially allows
// testing of error paths by forcing short reads.
type fixedReader struct {
	buf   []byte
	pos   int
	iobuf *bytes.Buffer
}

// Read reads the next len(p) bytes from the fixed reader.  When the number of
// bytes read would exceed the maximum number of allowed bytes to be read from
// the fixed writer, an error is returned.
//
// This satisfies the io.Reader interface.
func (fr *fixedReader) Read(p []byte) (n int, err error) {
	n, err = fr.iobuf.Read(p)
	fr.pos += n
	return
}

// newFixedReader returns a new io.Reader that will error once more bytes than
// the specified max have been read.
func newFixedReader(max int, buf []byte) io.Reader {
	b := make([]byte, max)
	if buf != nil {
		copy(b[:], buf)
	}

	iobuf := bytes.NewBuffer(b)
	fr := fixedReader{b, 0, iobuf}
	return &fr
}

func min(a, b uint32) uint32 {
    if a < b {
        return a
    }
    return b
}

func dcrdWire(pver uint32, net uint32, input []byte) {
    {
        rbuf := bytes.NewBuffer(input)
        var mblock dcrd_wire.MsgBlock
        mblock.DeserializeTxLoc(rbuf)
    }
    if len(input) >= 24 {
        length := binary.LittleEndian.Uint32(input[16:20])
        length = min(length, uint32(len(input) - 24))
        payload := input[24:24+length]
        checksum := dcrd_chainhash.HashB(payload)[0:4]
        copy(input[20:], checksum)
    }
    r := newFixedReader(len(input), input)
    _, msg, _, err := dcrd_wire.ReadMessageN(r, pver, dcrd_wire.CurrencyNet(net))
    if err != nil {
        return
    }
    switch msgx := msg.(type) {
        case *dcrd_wire.MsgVersion:
            break
		case *dcrd_wire.MsgVerAck:
            break
		case *dcrd_wire.MsgGetAddr:
            break
		case *dcrd_wire.MsgAddr:
            break
		case *dcrd_wire.MsgPing:
            break
		case *dcrd_wire.MsgPong:
            break
		case *dcrd_wire.MsgMemPool:
            break
		case *dcrd_wire.MsgGetMiningState:
            break
		case *dcrd_wire.MsgMiningState:
            break
		case *dcrd_wire.MsgTx:
            msgx.BytesPrefix()
            msgx.BytesWitness()
            msgx.PkScriptLocs()
            break
		case *dcrd_wire.MsgBlock:
            break
		case *dcrd_wire.MsgInv:
            break
		case *dcrd_wire.MsgHeaders:
            break
		case *dcrd_wire.MsgNotFound:
            break
		case *dcrd_wire.MsgGetData:
            break
        /* TODO */
		case *dcrd_wire.MsgCFHeaders:
            break
		case *dcrd_wire.MsgCFTypes:
            break
		case *dcrd_wire.MsgFeeFilter:
            break
		case *dcrd_wire.MsgReject:
            break
		case *dcrd_wire.MsgSendHeaders:
            break
    }
    var buf bytes.Buffer
    dcrd_wire.WriteMessageN(&buf, msg, pver, dcrd_wire.CurrencyNet(net))
}

func Fuzz(input []byte) {
    if len(input) < 8 {
        return
    }
    pver := binary.BigEndian.Uint32(input[0:4])
    net := binary.BigEndian.Uint32(input[4:8])
    input = input[8:]
    dcrdWire(pver, net, input)
}
