// Package cab allows you to interact with a Windows Cabinet File
package cab

import (
	"encoding/binary"
	"fmt"
	"os"
)

const cabSignature = "MSCF"

// type Cabinet is the main parser for Windows Cabinet Files
type Cabinet struct {
	// the types u1, u2,  and u4 are used to represent
	//  unsigned 8-, 16-, and 32-bit integer values
	// uint8_t[1] uint16_t[2] uint32_t[4]
	size uint32
}

// NewCabParser takes a filename location as string
func NewCabParser(fileName string) (*Cabinet, error) {
	var c Cabinet
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	header := make([]byte, 40)
	_, err = f.Read(header)
	if err != nil {
		return nil, err
	}
	sig := string(header[0:4])
	if sig != cabSignature {
		return nil, fmt.Errorf("%s signature: %s does not match '%s'",
			fileName, sig, cabSignature)
	}
	c.size = binary.LittleEndian.Uint32(header[8:12])
	fmt.Printf("sig: %s size: %+v rest: %#v\n", sig, c.size, header)
	return &c, nil
}
