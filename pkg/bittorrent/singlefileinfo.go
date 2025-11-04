package bittorrent

import (
	"crypto/sha1"
	"fmt"
	"github.com/gallyamow/go-bencoder"
)

// SingleFileInfo represents dictionary item for single file mode.
type SingleFileInfo struct {
	PieceLength int64      // number of bytes in each piece (integer)
	Piece       [][20]byte // string consisting of the concatenation of all 20-byte SHA1 hash values, one per piece (byte string, i.e. not urlencoded)
	Private     int        // (optional) this field is an integer. If it is set to "1", the client MUST publish its presence to get other peers ONLY via the trackers explicitly described in the metainfo file. If this field is set to "0" or is not present, the client may obtain peer from other means, e.g. PEX peer exchange, dht. Here, "private" may be read as "no external peer source".
	Name        string     // the filename. This is purely advisory. (string)
	Length      int64      // length of the file in bytes (integer)
	MD5sum      [32]byte   // (optional) a 32-character hexadecimal string corresponding to the MD5 sum of the file. This is not used by BitTorrent at all, but it is included by some programs for greater compatibility.
}

func (s SingleFileInfo) String() string {
	return fmt.Sprintf("SingleFileInfo{ PieceLength: %d, Piece: %v, Private: %d, Name: %s, Length: %d, MD5sum: %s }",
		s.PieceLength, s.Piece, s.Private, s.Name, s.Length, string(s.MD5sum[:]))
}

func (s SingleFileInfo) Size() int64 {
	return s.Length
}

func (s SingleFileInfo) Hash() [20]byte {
	return sha1.Sum(bencoder.Encode(s))
}

func (s SingleFileInfo) Parse(decoded map[string]any) {
	if val, ok := decoded["piece length"].(int64); ok {
		s.PieceLength = val
	}

	if val, ok := decoded["piece"].([][20]byte); ok {
		s.Piece = val
	}

	if val, ok := decoded["private"].(int); ok {
		s.Private = val
	}

	if val, ok := decoded["name"].(string); ok {
		s.Name = val
	}

	if val, ok := decoded["length"].(int64); ok {
		s.Length = val
	}

	if val, ok := decoded["md5sum"].([32]byte); ok {
		s.MD5sum = val
	}
}
