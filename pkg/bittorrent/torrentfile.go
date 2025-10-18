package bittorrent

import (
	"github.com/gallyamow/go-bencoder"
	"os"
	"reflect"
)

type TorrentFile[T SingleFileInfo | MultipleFileInfo] struct {
	Announce     string  // The announce URL of the tracker (string)
	CreationDate int64   // (optional) the creation time of the torrent, in standard UNIX epoch format (integer, seconds since 1-Jan-1970 00:00:00 UTC)
	Comment      *string // (optional) free-form textual comments of the author (string)
	CreatedBy    *string // (optional) name and version of the program used to create the .torrent (string)
	Encoding     *string // (optional) the string encoding format used to generate the pieces part of the info dictionary in the .torrent metafile (string)
	Info         *T
}

// SingleFileInfo represents dictionary item for single file mode
type SingleFileInfo struct {
	PieceLength int64      // number of bytes in each piece (integer)
	Piece       [][20]byte // string consisting of the concatenation of all 20-byte SHA1 hash values, one per piece (byte string, i.e. not urlencoded)
	Private     int        // (optional) this field is an integer. If it is set to "1", the client MUST publish its presence to get other peers ONLY via the trackers explicitly described in the metainfo file. If this field is set to "0" or is not present, the client may obtain peer from other means, e.g. PEX peer exchange, dht. Here, "private" may be read as "no external peer source".
	Name        string     // the filename. This is purely advisory. (string)
	Length      int64      // length of the file in bytes (integer)
	MD5sum      [32]byte   // (optional) a 32-character hexadecimal string corresponding to the MD5 sum of the file. This is not used by BitTorrent at all, but it is included by some programs for greater compatibility.
}

// MultipleFileInfo represents dictionary item for multiple file mode
type MultipleFileInfo struct {
	PieceLength int64      // number of bytes in each piece (integer)
	Piece       [][20]byte // string consisting of the concatenation of all 20-byte SHA1 hash values, one per piece (byte string, i.e. not urlencoded)
	Private     *int       // (optional) this field is an integer. If it is set to "1", the client MUST publish its presence to get other peers ONLY via the trackers explicitly described in the metainfo file. If this field is set to "0" or is not present, the client may obtain peer from other means, e.g. PEX peer exchange, dht. Here, "private" may be read as "no external peer source".
	Name        *string    // the filename. This is purely advisory. (string)
	Files       *MultipleFile
}

type MultipleFile struct {
	Length int64     // length of the file in bytes (integer)
	MD5sum *[32]byte // (optional) a 32-character hexadecimal string corresponding to the MD5 sum of the file. This is not used by BitTorrent at all, but it is included by some programs for greater compatibility.
	Path   []string  // a list containing one or more string elements that together represent the path and filename. Each element in the list corresponds to either a directory name or (in the case of the final element) the filename. For example, a the file "dir1/dir2/file.ext" would consist of three string elements: "dir1", "dir2", and "file.ext". This is encoded as a bencoded list of strings such as l4:dir14:dir28:file.exte
}

func OpenTorrentFile[T SingleFileInfo | MultipleFileInfo](path string) (*TorrentFile[T], error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	tmp, err := bencoder.Decode(file)
	if err != nil {
		return nil, err
	}

	torrentFile, err := parseTorrentFile[T](tmp.(map[string]any))
	if err != nil {
		return nil, err
	}

	return &torrentFile, nil
}

func parseTorrentFile[T SingleFileInfo | MultipleFileInfo](decoded map[string]any) (TorrentFile[T], error) {
	torrentFile := TorrentFile[T]{}

	if val, ok := decoded["creation date"].(int64); ok {
		torrentFile.CreationDate = val
	}

	if val, ok := decoded["announce"].(string); ok {
		torrentFile.Announce = val
	}

	if val, ok := decoded["comment"].(string); ok {
		torrentFile.Comment = &val
	}

	if val, ok := decoded["created by"].(string); ok {
		torrentFile.CreatedBy = &val
	}

	if val, ok := decoded["encoding"].(string); ok {
		torrentFile.Encoding = &val
	}

	var t T
	switch reflect.TypeOf(t) {
	case reflect.TypeOf(SingleFileInfo{}):
		info := parseSingleFileInfo(decoded["info"].(map[string]any))
		torrentFile.Info = any(&info).(*T)
	case reflect.TypeOf(MultipleFileInfo{}):
		panic("not implemented")
	}

	return torrentFile, nil
}

func parseSingleFileInfo(decoded map[string]any) SingleFileInfo {
	fileInfo := SingleFileInfo{}

	if val, ok := decoded["piece length"].(int64); ok {
		fileInfo.PieceLength = val
	}

	if val, ok := decoded["piece"].([][20]byte); ok {
		fileInfo.Piece = val
	}

	if val, ok := decoded["private"].(int); ok {
		fileInfo.Private = val
	}

	if val, ok := decoded["name"].(string); ok {
		fileInfo.Name = val
	}

	if val, ok := decoded["length"].(int64); ok {
		fileInfo.Length = val
	}

	if val, ok := decoded["md5sum"].([32]byte); ok {
		fileInfo.MD5sum = val
	}

	return fileInfo
}
