package bittorrent

import (
	"fmt"
)

type MultipleFile struct {
	Length int64     // length of the file in bytes (integer)
	MD5sum *[32]byte // (optional) a 32-character hexadecimal string corresponding to the MD5 sum of the file. This is not used by BitTorrent at all, but it is included by some programs for greater compatibility.
	Path   []string  // a list containing one or more string elements that together represent the path and filename. Each element in the list corresponds to either a directory name or (in the case of the final element) the filename. For example, a the file "dir1/dir2/file.ext" would consist of three string elements: "dir1", "dir2", and "file.ext". This is encoded as a bencoded list of strings such as l4:dir14:dir28:file.exte
}

func (m *MultipleFile) String() string {
	return fmt.Sprintf("MultipleFile{ Length: %d, MD5sum: %x, Path: %q }", m.Length, m.MD5sum, m.Path)
}

// MultipleFileInfo represents dictionary item for multiple file mode
type MultipleFileInfo struct {
	PieceLength int64      // number of bytes in each piece (integer)
	Piece       [][20]byte // string consisting of the concatenation of all 20-byte SHA1 hash values, one per piece (byte string, i.e. not urlencoded)
	Private     *int       // (optional) this field is an integer. If it is set to "1", the client MUST publish its presence to get other peers ONLY via the trackers explicitly described in the metainfo file. If this field is set to "0" or is not present, the client may obtain peer from other means, e.g. PEX peer exchange, dht. Here, "private" may be read as "no external peer source".
	Name        *string    // the filename. This is purely advisory. (string)
	Files       *[]MultipleFile
}

func (m MultipleFileInfo) String() string {
	return fmt.Sprintf("MultipleFileInfo{ PieceLength: %d, Piece: %x, Private: %v, Name: %q, Files: %v }", m.PieceLength, m.Piece, m.Private, StringifyPtr(m.Name), StringifyPtr(m.Files))
}

func (m MultipleFileInfo) Size() int64 {
	size := int64(0)
	for _, file := range *m.Files {
		size += file.Length
	}
	return size
}

func (m MultipleFileInfo) Hash() [32]byte {
	panic("not implemented")
}

func (m MultipleFileInfo) Parse(decoded map[string]any) {
	panic("not implemented")
}
