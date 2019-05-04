// Package cab allows you to interact with a Windows Cabinet File
package cab

// type Cabinet is the main parser for Windows Cabinet Files
type Cabinet struct {
	// the types u1, u2,  and u4 are used to represent
	//  unsigned 8-, 16-, and 32-bit integer values
	setID int16
}

// NewCabParser takes a filename location as string
func NewCabParser(f string) Cabinet {
	return nil
}
