package chain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
)

func Sha256(block *Block) Hash {
	s := sha256.New()

	byteWriter := bytes.NewBuffer([]byte{})

	byteWriter.WriteString(string(block.Nonce))
	byteWriter.WriteString(strconv.Itoa(block.Timestamp.Nanosecond()))

	if block.Previous != nil {
		byteWriter.WriteString(string(block.Previous.Hash))
	}

	io.Copy(s, byteWriter)

	return Hash(fmt.Sprintf("%x", s.Sum(nil)))
}
