// Package cab allows you to interact with a Windows Cabinet File
package cab

import (
	"encoding/binary"
	"fmt"
	"os"
)

const (
	cabSignature  = "MSCF"
	HDRPREVCAB    = 0x0001
	HDRNEXTCAB    = 0x0002
	HDRRESPRESENT = 0x0004
)

// type Cabinet is the main parser for Windows Cabinet Files
type Cabinet struct {
	// the types u1, u2,  and u4 are used to represent
	//  unsigned 8-, 16-, and 32-bit integer values
	// uint8_t[1] uint16_t[2] uint32_t[4]
	Size             uint32 // Full size of cabinet size in bytes
	VersionMajor     uint8  // Major Verison (should be 1)
	VersionMinor     uint8  // Minor Version (should be 3)
	Flags            uint16 // Header flags
	SetID            uint16
	SequencePosition uint16
}

// NewCabinet takes a filename location as string
func NewCabinet(fileName string) (*Cabinet, error) {
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
	c.Size = binary.LittleEndian.Uint32(header[8:12])
	cabFileOffset := binary.LittleEndian.Uint32(header[16:20])
	c.VersionMinor = header[24]
	c.VersionMajor = header[25]
	folderCount := binary.LittleEndian.Uint16(header[26:28])
	fileCount := binary.LittleEndian.Uint16(header[29:31])
	c.Flags = binary.LittleEndian.Uint16(header[32:34])
	c.SetID = binary.LittleEndian.Uint16(header[35:37])
	c.SequencePosition = binary.LittleEndian.Uint16(header[38:40])
	switch {
	case c.Flags&HDRPREVCAB != 0:
		fmt.Println("Previous cabinet present")
	case c.Flags&HDRNEXTCAB != 0:
		fmt.Println("Next cabinet present")
	case c.Flags&HDRRESPRESENT != 0:
		fmt.Println("Reserve present")
	}
	fmt.Printf("sig: %s offset: %+v foldercount: %+v filecount: %+v Cabinet: %+v\n",
		sig, cabFileOffset, folderCount, fileCount, c)
	return &c, nil
}
